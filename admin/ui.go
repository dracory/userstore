package admin

import (
	"errors"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/userstore/admin/shared"
	adminUsers "github.com/gouniverse/userstore/admin/users"
	"github.com/gouniverse/utils"
)

func UI(config shared.Config) (hb.TagInterface, error) {
	if config.ResponseWriter == nil {
		return nil, errors.New("ResponseWriter is required")
	}

	if config.Request == nil {
		return nil, errors.New("Request is required")
	}

	if config.Store == nil {
		return nil, errors.New("Store is required")
	}

	if config.Logger == nil {
		return nil, errors.New("Logger is required")
	}

	if config.Layout == nil {
		return nil, errors.New("Layout is required")
	}

	if len(config.TokenizedColumns) > 0 {
		if config.TokensBulk == nil {
			return nil, errors.New("TokensBulk function is required")
		}

		if config.TokensRead == nil {
			return nil, errors.New("TokensRead function is required")
		}
	}

	return handler(config), nil
}

func handler(config shared.Config) hb.TagInterface {
	controller := utils.Req(config.Request, "controller", "")

	if controller == "" {
		controller = shared.PathHome
	}

	if controller == shared.PathHome {
		return NewHomeController().ToTag(config)
	}

	if controller == shared.PathUserCreate {
		return adminUsers.NewUserCreateController().ToTag(config)
	}

	if controller == shared.PathUserDelete {
		return adminUsers.NewUserDeleteController().ToTag(config)
	}

	if controller == shared.PathUserUpdate {
		return adminUsers.NewUserUpdateController().ToTag(config)
	}

	if controller == shared.PathUsers {
		return adminUsers.NewUserManagerController().ToTag(config)
	}

	html := config.Layout(config.ResponseWriter, config.Request, shared.LayoutOptions{
		Title: "Path not found",
		Body:  hb.H1().HTML(controller).ToHTML(),
	})

	return hb.Raw(html)
}
