package admin

import (
	"context"

	"github.com/dracory/req"
	"github.com/dracory/userstore"
	"github.com/dracory/userstore/admin/shared"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
)

type userDeleteController struct{}

var _ shared.PageInterface = (*userDeleteController)(nil)

type userDeleteControllerData struct {
	config         shared.Config
	userID         string
	user           userstore.UserInterface
	successMessage string
	//errorMessage   string
}

func NewUserDeleteController() *userDeleteController {
	return &userDeleteController{}
}

func (controller userDeleteController) ToTag(config shared.Config) hb.TagInterface {
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

	return controller.modal(data)
}

func (controller *userDeleteController) modal(data userDeleteControllerData) hb.TagInterface {
	submitUrl := shared.Url(data.config.Request, shared.PathUserDelete, map[string]string{
		"user_id": data.userID,
	})

	modalID := "ModalUserDelete"
	modalBackdropClass := "ModalBackdrop"

	formGroupUserId := hb.Input().
		Type(hb.TYPE_HIDDEN).
		Name("user_id").
		Value(data.userID)

	buttonDelete := hb.Button().
		HTML("Delete").
		Class("btn btn-primary float-end").
		HxInclude("#Modal" + modalID).
		HxPost(submitUrl).
		HxSelectOob("#ModalUserDelete").
		HxTarget("body").
		HxSwap("beforeend")

	modalCloseScript := `closeModal` + modalID + `();`

	modalHeading := hb.Heading5().HTML("Delete User").Style(`margin:0px;`)

	modalClose := hb.Button().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	jsCloseFn := `function closeModal` + modalID + `() {document.getElementById('ModalUserDelete').remove();[...document.getElementsByClassName('` + modalBackdropClass + `')].forEach(el => el.remove());}`

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
						Child(hb.Paragraph().Text("Are you sure you want to delete this user?").Style(`margin-bottom:20px;color:red;`)).
						Child(hb.Paragraph().Text("This action cannot be undone.")).
						Child(formGroupUserId)).
				Child(bs.ModalFooter().
					Style(`display:flex;justify-content:space-between;`).
					Child(
						hb.Button().HTML("Close").
							Class("btn btn-secondary float-start").
							Data("bs-dismiss", "modal").
							OnClick(modalCloseScript)).
					Child(buttonDelete)),
			))

	backdrop := hb.Div().Class(modalBackdropClass).
		Class("modal-backdrop fade show").
		Style("display:block;z-index:1000;")

	return hb.Wrap().
		Children([]hb.TagInterface{
			modal,
			backdrop,
		})
}

func (controller *userDeleteController) prepareDataAndValidate(config shared.Config) (data userDeleteControllerData, errorMessage string) {
	data.config = config
	data.userID = req.GetString(config.Request, "user_id")

	if data.userID == "" {
		return data, "user id is required"
	}

	user, err := config.Store.UserFindByID(context.Background(), data.userID)

	if err != nil {
		config.Logger.Error("Error. At userDeleteController > prepareDataAndValidate", "error", err.Error())
		return data, "User not found"
	}

	if user == nil {
		return data, "User not found"
	}

	data.user = user

	if config.Request.Method != "POST" {
		return data, ""
	}

	err = config.Store.UserSoftDelete(context.Background(), user)

	if err != nil {
		config.Logger.Error("Error. At userDeleteController > prepareDataAndValidate", "error", err.Error())
		return data, "Deleting user failed. Please contact an administrator."
	}

	data.successMessage = "user deleted successfully."

	return data, ""

}
