package admin

import (
	"context"
	"net/http"
	"strings"

	"github.com/dracory/form"
	"github.com/dracory/req"
	"github.com/dracory/sb"
	"github.com/dracory/userstore"
	"github.com/dracory/userstore/admin/shared"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

const ActionModalUserFilterShow = "modal_user_filter_show"

// == CONTROLLER ==============================================================

type userManagerController struct{}

var _ shared.PageInterface = (*userManagerController)(nil)

// == CONSTRUCTOR =============================================================

func NewUserManagerController() *userManagerController {
	return &userManagerController{}
}

func (c *userManagerController) ToTag(config shared.Config) hb.TagInterface {
	html, withLayout := c.checkAndProcess(config)

	if !withLayout {
		return hb.Raw(html)
	}

	layout := config.Layout(config.ResponseWriter, config.Request, shared.LayoutOptions{
		Title: `Users | User Manager`,
		Body:  html,
		Scripts: []string{
			shared.ScriptHtmx,
			shared.ScriptSwal,
		},
	})

	return hb.Raw(layout)
}

func (controller *userManagerController) checkAndProcess(config shared.Config) (html string, withLayout bool) {
	data, errorMessage := controller.prepareData(config)

	if errorMessage != "" {
		return hb.Div().
			Class("alert alert-danger").
			Text(errorMessage).
			ToHTML(), true
	}

	if data.action == ActionModalUserFilterShow {
		return controller.onModalUserFilterShow(data).ToHTML(), false
	}

	return controller.page(data).ToHTML(), true
}

func (controller *userManagerController) onModalUserFilterShow(data userManagerControllerData) *hb.Tag {
	modalCloseScript := `document.getElementById('ModalMessage').remove();document.getElementById('ModalBackdrop').remove();`

	title := hb.Heading5().
		Text("Filters").
		Style(`margin:0px;padding:0px;`)

	buttonModalClose := hb.Button().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	buttonCancel := hb.Button().
		Child(hb.I().Class("bi bi-chevron-left me-2")).
		HTML("Cancel").
		Class("btn btn-secondary float-start").
		OnClick(modalCloseScript)

	buttonOk := hb.Button().
		Child(hb.I().Class("bi bi-check me-2")).
		HTML("Apply").
		Class("btn btn-primary float-end").
		OnClick(`FormFilters.submit();` + modalCloseScript)

	filterForm := form.NewForm(form.FormOptions{
		ID:     "FormFilters",
		Method: http.MethodGet,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Label: "Status",
				Name:  "status",
				Type:  form.FORM_FIELD_TYPE_SELECT,
				Help:  `The status of the user.`,
				Value: data.formStatus,
				Options: []form.FieldOption{
					{
						Value: "",
						Key:   "",
					},
					{
						Value: "Active",
						Key:   userstore.USER_STATUS_ACTIVE,
					},
					{
						Value: "Inactive",
						Key:   userstore.USER_STATUS_INACTIVE,
					},
					{
						Value: "Unverified",
						Key:   userstore.USER_STATUS_UNVERIFIED,
					},
					{
						Value: "Deleted",
						Key:   userstore.USER_STATUS_DELETED,
					},
				},
			}),
			form.NewField(form.FieldOptions{
				Label: "First Name",
				Name:  "first_name",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formFirstName,
				Help:  `Filter by first name.`,
			}),
			form.NewField(form.FieldOptions{
				Label: "Last Name",
				Name:  "last_name",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formLastName,
				Help:  `Filter by last name.`,
			}),
			form.NewField(form.FieldOptions{
				Label: "Email",
				Name:  "email",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formEmail,
				Help:  `Filter by email.`,
			}),
			form.NewField(form.FieldOptions{
				Label: "Created From",
				Name:  "created_from",
				Type:  form.FORM_FIELD_TYPE_DATE,
				Value: data.formCreatedFrom,
				Help:  `Filter by creation date.`,
			}),
			form.NewField(form.FieldOptions{
				Label: "Created To",
				Name:  "created_to",
				Type:  form.FORM_FIELD_TYPE_DATE,
				Value: data.formCreatedTo,
				Help:  `Filter by creation date.`,
			}),
			form.NewField(form.FieldOptions{
				Label: "User ID",
				Name:  "user_id",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formUserID,
				Help:  `Find user by reference number (ID).`,
			}),
		},
	}).Build()

	modal := bs.Modal().
		ID("ModalMessage").
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Children([]hb.TagInterface{
			bs.ModalDialog().Children([]hb.TagInterface{
				bs.ModalContent().Children([]hb.TagInterface{
					bs.ModalHeader().Children([]hb.TagInterface{
						title,
						buttonModalClose,
					}),

					bs.ModalBody().
						Child(filterForm),

					bs.ModalFooter().
						Style(`display:flex;justify-content:space-between;`).
						Child(buttonCancel).
						Child(buttonOk),
				}),
			}),
		})

	backdrop := hb.Div().
		ID("ModalBackdrop").
		Class("modal-backdrop fade show").
		Style("display:block;")

	return hb.Wrap().Children([]hb.TagInterface{
		modal,
		backdrop,
	})

}

func (controller *userManagerController) page(data userManagerControllerData) hb.TagInterface {
	breadcrumbs := shared.Breadcrumbs(data.config, []shared.Breadcrumb{
		{
			Name: "User Manager",
			URL:  shared.Url(data.config.Request, shared.PathUsers, nil),
		},
	})

	buttonUserNew := hb.Button().
		Class("btn btn-primary float-end").
		Child(hb.I().Class("bi bi-plus-circle").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
		HTML("New User").
		HxGet(shared.Url(data.config.Request, shared.PathUserCreate, nil)).
		HxTarget("body").
		HxSwap("beforeend")

	title := hb.Heading1().
		HTML("Users. User Manager").
		Child(buttonUserNew)

	return hb.Div().
		Class("container").
		Child(breadcrumbs).
		Child(hb.HR()).
		Child(title).
		Child(controller.tableUsers(data))
}

func (controller *userManagerController) tableUsers(data userManagerControllerData) hb.TagInterface {
	table := hb.Table().
		Class("table table-striped table-hover table-bordered").
		Children([]hb.TagInterface{
			hb.Thead().Children([]hb.TagInterface{
				hb.TR().Children([]hb.TagInterface{
					hb.TH().
						Child(controller.sortableColumnLabel(data, "First Name", "first_name")).
						Text(", ").
						Child(controller.sortableColumnLabel(data, "Last Name", "last_name")).
						Text(", ").
						Child(controller.sortableColumnLabel(data, "Reference", "id")).
						Style(`cursor: pointer;`),
					hb.TH().
						Child(controller.sortableColumnLabel(data, "Status", "status")).
						Style("width: 200px;cursor: pointer;"),
					hb.TH().
						Child(controller.sortableColumnLabel(data, "E-mail", "email")).
						Style("width: 1px;cursor: pointer;"),
					hb.TH().
						Child(controller.sortableColumnLabel(data, "Created", "created_at")).
						Style("width: 1px;cursor: pointer;"),
					hb.TH().
						Child(controller.sortableColumnLabel(data, "Modified", "updated_at")).
						Style("width: 1px;cursor: pointer;"),
					hb.TH().
						HTML("Actions"),
				}),
			}),
			hb.Tbody().Children(lo.Map(data.userList, func(user userstore.UserInterface, _ int) hb.TagInterface {
				// firstName, lastName, email, err := userUntokenize(data.config, user)
				firstName := user.FirstName()
				lastName := user.LastName()
				email := user.Email()

				untokenized, err := userUntokenize(data.config, user)

				if err != nil {
					data.config.Logger.Error("At userManagerController > tableUsers", "error", err.Error())
					firstName = "n/a"
					lastName = "n/a"
					email = "n/a"
				}

				if lo.HasKey(untokenized, userstore.COLUMN_FIRST_NAME) {
					firstName = untokenized[userstore.COLUMN_FIRST_NAME]
				}

				if lo.HasKey(untokenized, userstore.COLUMN_LAST_NAME) {
					lastName = untokenized[userstore.COLUMN_LAST_NAME]
				}

				if lo.HasKey(untokenized, userstore.COLUMN_EMAIL) {
					email = untokenized[userstore.COLUMN_EMAIL]
				}

				userLink := hb.Hyperlink().
					Text(firstName).
					Text(` `).
					Text(lastName).
					Href(shared.Url(data.config.Request, shared.PathUserUpdate, map[string]string{"user_id": user.ID()}))

				status := hb.Span().
					Style(`font-weight: bold;`).
					StyleIf(user.IsActive(), `color:green;`).
					StyleIf(user.IsSoftDeleted(), `color:silver;`).
					StyleIf(user.IsUnverified(), `color:blue;`).
					StyleIf(user.IsInactive(), `color:red;`).
					HTML(user.Status())

				buttonEdit := hb.Hyperlink().
					Class("btn btn-primary me-2").
					Child(hb.I().Class("bi bi-pencil-square")).
					Title("Edit").
					Href(shared.Url(data.config.Request, shared.PathUserUpdate, map[string]string{"user_id": user.ID()})).
					Target("_blank")

				buttonDelete := hb.Hyperlink().
					Class("btn btn-danger").
					Child(hb.I().Class("bi bi-trash")).
					Title("Delete").
					HxGet(shared.Url(data.config.Request, shared.PathUserDelete, map[string]string{"user_id": user.ID()})).
					HxTarget("body").
					HxSwap("beforeend")

				buttonImpersonate := hb.Hyperlink().
					Class("btn btn-warning me-2").
					Child(hb.I().Class("bi bi-shuffle")).
					Title("Impersonate").
					Href(shared.Url(data.config.Request, shared.PathUserImpersonate, map[string]string{"user_id": user.ID()}))

				return hb.TR().Children([]hb.TagInterface{
					hb.TD().
						Child(hb.Div().Child(userLink)).
						Child(hb.Div().
							Style("font-size: 11px;").
							HTML("Ref: ").
							HTML(user.ID())),
					hb.TD().
						Child(status),
					hb.TD().
						Child(hb.Div().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(email)),
					hb.TD().
						Child(hb.Div().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(user.CreatedAtCarbon().Format("d M Y"))),
					hb.TD().
						Child(hb.Div().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(user.UpdatedAtCarbon().Format("d M Y"))),
					hb.TD().
						Child(buttonEdit).
						Child(buttonImpersonate).
						Child(buttonDelete),
				})
			})),
		})

	// cfmt.Successln("Table: ", table)

	return hb.Wrap().Children([]hb.TagInterface{
		controller.tableFilter(data),
		table,
		controller.tablePagination(data, int(data.userCount), data.pageInt, data.perPage),
	})
}

func (controller *userManagerController) sortableColumnLabel(
	data userManagerControllerData,
	tableLabel string,
	columnName string,
) hb.TagInterface {
	isSelected := strings.EqualFold(data.sortBy, columnName)

	direction := lo.If(data.sortOrder == "asc", "desc").Else("asc")

	if !isSelected {
		direction = "asc"
	}

	link := shared.Url(data.config.Request, shared.PathUsers, map[string]string{
		"page":      "0",
		"by":        columnName,
		"sort":      direction,
		"date_from": data.formCreatedFrom,
		"date_to":   data.formCreatedTo,
		"status":    data.formStatus,
		"user_id":   data.formUserID,
	})
	return hb.Hyperlink().
		HTML(tableLabel).
		Child(controller.sortingIndicator(columnName, data.sortBy, direction)).
		Href(link)
}

func (controller *userManagerController) sortingIndicator(
	columnName,
	sortByColumnName,
	sortOrder string,
) hb.TagInterface {
	isSelected := strings.EqualFold(sortByColumnName, columnName)

	direction := lo.If(isSelected && sortOrder == "asc", "up").
		ElseIf(isSelected && sortOrder == "desc", "down").
		Else("none")

	sortingIndicator := hb.Span().
		Class("sorting").
		HTMLIf(direction == "up", "&#8595;").
		HTMLIf(direction == "down", "&#8593;").
		HTMLIf(direction != "down" && direction != "up", "")

	return sortingIndicator
}

func (controller *userManagerController) tableFilter(data userManagerControllerData) hb.TagInterface {
	buttonFilter := hb.Button().
		Class("btn btn-sm btn-info me-2").
		Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
		Child(hb.I().Class("bi bi-filter me-2")).
		Text("Filters").
		HxPost(shared.Url(data.config.Request, shared.PathUsers, map[string]string{
			"action":       ActionModalUserFilterShow,
			"first_name":   data.formFirstName,
			"last_name":    data.formLastName,
			"email":        data.formEmail,
			"status":       data.formStatus,
			"user_id":      data.formUserID,
			"created_from": data.formCreatedFrom,
			"created_to":   data.formCreatedTo,
		})).
		HxTarget("body").
		HxSwap("beforeend")

	description := []string{
		hb.Span().HTML("Showing users").Text(" ").ToHTML(),
	}

	if data.formStatus != "" {
		description = append(description, hb.Span().Text("with status: "+data.formStatus).ToHTML())
	} else {
		description = append(description, hb.Span().Text("with status: any").ToHTML())
	}

	if data.formEmail != "" {
		description = append(description, hb.Span().Text("and email: "+data.formEmail).ToHTML())
	}

	if data.formUserID != "" {
		description = append(description, hb.Span().Text("and ID: "+data.formUserID).ToHTML())
	}

	if data.formFirstName != "" {
		description = append(description, hb.Span().Text("and first name: "+data.formFirstName).ToHTML())
	}

	if data.formLastName != "" {
		description = append(description, hb.Span().Text("and last name: "+data.formLastName).ToHTML())
	}

	if data.formCreatedFrom != "" && data.formCreatedTo != "" {
		description = append(description, hb.Span().Text("and created between: "+data.formCreatedFrom+" and "+data.formCreatedTo).ToHTML())
	} else if data.formCreatedFrom != "" {
		description = append(description, hb.Span().Text("and created after: "+data.formCreatedFrom).ToHTML())
	} else if data.formCreatedTo != "" {
		description = append(description, hb.Span().Text("and created before: "+data.formCreatedTo).ToHTML())
	}

	return hb.Div().
		Class("card bg-light mb-3").
		Style("").
		Children([]hb.TagInterface{
			hb.Div().Class("card-body").
				Child(buttonFilter).
				Child(hb.Span().
					HTML(strings.Join(description, " "))),
		})
}

func (controller *userManagerController) tablePagination(data userManagerControllerData, count, page, perPage int) hb.TagInterface {
	url := shared.Url(data.config.Request, shared.PathUsers, map[string]string{
		"status":       data.formStatus,
		"first_name":   data.formFirstName,
		"last_name":    data.formLastName,
		"email":        data.formEmail,
		"created_from": data.formCreatedFrom,
		"created_to":   data.formCreatedTo,
		"by":           data.sortBy,
		"order":        data.sortOrder,
	})

	url = lo.Ternary(strings.Contains(url, "?"), url+"&page=", url+"?page=") // page must be last

	pagination := bs.Pagination(bs.PaginationOptions{
		NumberItems:       count,
		CurrentPageNumber: page,
		PagesToShow:       5,
		PerPage:           perPage,
		URL:               url,
	})

	return hb.Div().
		Class(`d-flex justify-content-left mt-5 pagination-primary-soft rounded mb-0`).
		HTML(pagination)
}

func (controller *userManagerController) prepareData(config shared.Config) (data userManagerControllerData, errorMessage string) {
	var err error
	data.config = config
	data.action = req.GetStringTrimmed(config.Request, "action")
	data.page = req.GetStringTrimmedOr(config.Request, "page", "0")
	data.pageInt = cast.ToInt(data.page)
	data.perPage = cast.ToInt(req.GetStringTrimmedOr(config.Request, "per_page", "10"))
	data.sortOrder = req.GetStringTrimmedOr(config.Request, "sort_order", sb.DESC)
	data.sortBy = req.GetStringTrimmedOr(config.Request, "by", userstore.COLUMN_CREATED_AT)
	data.formEmail = req.GetStringTrimmed(config.Request, "email")
	data.formFirstName = req.GetStringTrimmed(config.Request, "first_name")
	data.formLastName = req.GetStringTrimmed(config.Request, "last_name")
	data.formStatus = req.GetStringTrimmed(config.Request, "status")
	data.formCreatedFrom = req.GetStringTrimmed(config.Request, "created_from")
	data.formCreatedTo = req.GetStringTrimmed(config.Request, "created_to")

	userList, userCount, err := controller.fetchUserList(data)

	if err != nil {
		config.Logger.Error("At userManagerController > prepareData", "error", err.Error())
		return data, "error retrieving users"
	}

	data.userList = userList
	data.userCount = userCount

	return data, ""
}

func (controller *userManagerController) fetchUserList(data userManagerControllerData) (users []userstore.UserInterface, userCount int64, err error) {
	userIDs := []string{}

	// if data.formFirstName != "" {
	// 	firstNameUserIDs, err := config.BlindIndexStoreFirstName.Search(data.formFirstName, blindindexstore.SEARCH_TYPE_CONTAINS)

	// 	if err != nil {
	// 		data.config.Logger.Error("At userManagerController > prepareData", err.Error())
	// 		return []userstore.UserInterface{}, 0, err
	// 	}

	// 	if len(firstNameUserIDs) == 0 {
	// 		return []userstore.UserInterface{}, 0, nil
	// 	}

	// 	userIDs = append(userIDs, firstNameUserIDs...)
	// }

	// if data.formLastName != "" {
	// 	lastNameUserIDs, err := config.BlindIndexStoreLastName.Search(data.formLastName, blindindexstore.SEARCH_TYPE_CONTAINS)

	// 	if err != nil {
	// 		config.Logger.Error("At userManagerController > prepareData", err.Error())
	// 		return []userstore.UserInterface{}, 0, err
	// 	}

	// 	if len(lastNameUserIDs) == 0 {
	// 		return []userstore.UserInterface{}, 0, nil
	// 	}

	// 	userIDs = append(userIDs, lastNameUserIDs...)
	// }

	// if data.formEmail != "" {
	// 	emailUserIDs, err := config.BlindIndexStoreEmail.Search(data.formEmail, blindindexstore.SEARCH_TYPE_CONTAINS)

	// 	if err != nil {
	// 		config.Logger.Error("At userManagerController > prepareData", err.Error())
	// 		return []userstore.UserInterface{}, 0, err
	// 	}

	// 	if len(emailUserIDs) == 0 {
	// 		return []userstore.UserInterface{}, 0, nil
	// 	}

	// 	userIDs = append(userIDs, emailUserIDs...)
	// }

	// query := userstore.UserQueryOptions{
	// 	IDIn:      userIDs,
	// 	Offset:    data.pageInt * data.perPage,
	// 	Limit:     data.perPage,
	// 	Status:    data.formStatus,
	// 	SortOrder: data.sortOrder,
	// 	OrderBy:   data.sortBy,
	// }

	query := userstore.NewUserQuery()

	if len(userIDs) > 0 {
		query = query.SetIDIn(userIDs)
	}

	if data.formStatus != "" {
		query = query.SetStatus(data.formStatus)
	}

	query = query.SetSortDirection(data.sortOrder)

	query = query.SetOrderBy(data.sortBy)

	query = query.SetOffset(data.pageInt * data.perPage)

	query = query.SetLimit(data.perPage)

	if data.formCreatedFrom != "" {
		query = query.SetCreatedAtGte(data.formCreatedFrom + " 00:00:00")
	}

	if data.formCreatedTo != "" {
		query = query.SetCreatedAtLte(data.formCreatedTo + " 23:59:59")
	}

	userList, err := data.config.Store.UserList(context.Background(), query)

	if err != nil {
		data.config.Logger.Error("At userManagerController > prepareData", "error", err.Error())
		return []userstore.UserInterface{}, 0, err
	}

	userCount, err = data.config.Store.UserCount(context.Background(), query)

	if err != nil {
		data.config.Logger.Error("At userManagerController > prepareData", "error", err.Error())
		return []userstore.UserInterface{}, 0, err
	}

	return userList, userCount, nil
}

type userManagerControllerData struct {
	config          shared.Config
	action          string
	page            string
	pageInt         int
	perPage         int
	sortOrder       string
	sortBy          string
	formStatus      string
	formEmail       string
	formFirstName   string
	formLastName    string
	formCreatedFrom string
	formCreatedTo   string
	formUserID      string
	userList        []userstore.UserInterface
	userCount       int64
}
