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

	if len(config.Tokenized) > 0 {
		if config.Tokenize == nil {
			return nil, errors.New("Tokenize function is required")
		}

		if config.Untokenize == nil {
			return nil, errors.New("Untokenize function is required")
		}
	} else {
		config.Tokenize = func(clear []string) (tokens []string) { return clear }
		config.Untokenize = func(tokens []string) (clear []string) { return tokens }
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

	if controller == shared.PathUsers {
		return adminUsers.NewUserManagerController().ToTag(config)
	}

	if controller == shared.PathUserCreate {
		return adminUsers.NewUserCreateController().ToTag(config)
	}

	html := config.Layout(config.ResponseWriter, config.Request, shared.LayoutOptions{
		Title: "Path not found",
		Body:  hb.H1().HTML(controller).ToHTML(),
	})

	return hb.Raw(html)
}
