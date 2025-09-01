package admin

import (
	"context"
	"net/http"

	"github.com/dracory/bs"
	"github.com/dracory/req"
	"github.com/dracory/userstore"
	"github.com/dracory/userstore/admin/shared"
	"github.com/gouniverse/hb"
)

type userCreateController struct{}

var _ shared.PageInterface = (*userCreateController)(nil)

type userCreateControllerData struct {
	config         shared.Config
	firstName      string
	lastName       string
	email          string
	successMessage string
	//errorMessage   string
}

func NewUserCreateController() *userCreateController {
	return &userCreateController{}
}

func (controller userCreateController) ToTag(config shared.Config) hb.TagInterface {
	data, errorMessage := controller.prepareDataAndValidate(config)

	if errorMessage != "" {
		return hb.Swal(hb.SwalOptions{
			Icon: "error",
			Text: errorMessage,
		})
	}

	if data.successMessage != "" {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon: "success",
				Text: data.successMessage,
			})).
			Child(hb.Script("setTimeout(() => {window.location.href = window.location.href}, 2000)"))
	}

	return controller.
		modal(data)
}

func (controller *userCreateController) modal(data userCreateControllerData) hb.TagInterface {
	submitUrl := shared.Url(data.config.Request, shared.PathUserCreate, map[string]string{})

	formGroupFirstName := bs.FormGroup().
		Class("mb-3").
		Child(bs.FormLabel("First name")).
		Child(bs.FormInput().Name("user_first_name").Value(data.firstName))

	formGroupLastName := bs.FormGroup().
		Class("mb-3").
		Child(bs.FormLabel("Last name")).
		Child(bs.FormInput().Name("user_last_name").Value(data.lastName))

	formGroupEmail := bs.FormGroup().
		Class("mb-3").
		Child(bs.FormLabel("Email")).
		Child(bs.FormInput().Name("user_email").Value(data.email))

	modalID := "ModaluserCreate"
	modalBackdropClass := "ModalBackdrop"

	modalCloseScript := `closeModal` + modalID + `();`

	modalHeading := hb.Heading5().HTML("New user Create").Style(`margin:0px;`)

	modalClose := hb.Button().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	jsCloseFn := `function closeModal` + modalID + `() {document.getElementById('ModaluserCreate').remove();[...document.getElementsByClassName('` + modalBackdropClass + `')].forEach(el => el.remove());}`

	buttonSend := hb.Button().
		Child(hb.I().Class("bi bi-check me-2")).
		HTML("Create & Edit").
		Class("btn btn-primary float-end").
		HxInclude("#" + modalID).
		HxPost(submitUrl).
		HxSelectOob("#ModaluserCreate").
		HxTarget("body").
		HxSwap("beforeend")

	buttonCancel := hb.Button().
		Child(hb.I().Class("bi bi-chevron-left me-2")).
		HTML("Close").
		Class("btn btn-secondary float-start").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	modal := bs.Modal().
		ID(modalID).
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Child(hb.Script(jsCloseFn)).
		Child(bs.ModalDialog().
			Child(bs.ModalContent().
				Child(
					bs.ModalHeader().
						Child(modalHeading).
						Child(modalClose)).
				Child(
					bs.ModalBody().
						Child(formGroupFirstName).
						Child(formGroupLastName).
						Child(formGroupEmail)).
				Child(bs.ModalFooter().
					Style(`display:flex;justify-content:space-between;`).
					Child(buttonCancel).
					Child(buttonSend)),
			))

	backdrop := hb.Div().Class(modalBackdropClass).
		Class("modal-backdrop fade show").
		Style("display:block;z-index:1000;")

	return hb.Wrap().Children([]hb.TagInterface{
		modal,
		backdrop,
	})
}

func (controller *userCreateController) prepareDataAndValidate(config shared.Config) (data userCreateControllerData, errorMessage string) {
	data.config = config
	data.firstName = req.GetStringTrimmed(config.Request, "user_first_name")
	data.lastName = req.GetStringTrimmed(config.Request, "user_last_name")
	data.email = req.GetStringTrimmed(config.Request, "user_email")

	if config.Request.Method != http.MethodPost {
		return data, ""
	}

	if data.firstName == "" {
		return data, "user first name is required"
	}

	if data.lastName == "" {
		return data, "user last name is required"
	}

	if data.email == "" {
		return data, "user email is required"
	}

	user := userstore.NewUser()
	user.SetFirstName(data.firstName)
	user.SetLastName(data.lastName)
	user.SetEmail(data.email)

	err := config.Store.UserCreate(context.Background(), user)

	if err != nil {
		config.Logger.Error("Error. At userCreateController > prepareDataAndValidate", "error", err.Error())
		return data, "Creating user failed. Please contact an administrator."
	}

	data.successMessage = "user created successfully."

	return data, ""

}
