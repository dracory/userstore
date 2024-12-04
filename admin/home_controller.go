package admin

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/userstore/admin/shared"
)

// == CONSTRUCTOR ==============================================================

func NewHomeController() shared.PageInterface {
	return &homeController{}
}

// == CONTROLLER ===============================================================

type homeController struct{}

type homeControllerData struct {
	shared.Config
}

func (c *homeController) ToTag(config shared.Config) hb.TagInterface {
	html := c.checkAndProcess(config)

	layout := config.Layout(config.ResponseWriter, config.Request, shared.LayoutOptions{
		Title: `Dashboard | Users`,
		Body:  html,
		Scripts: []string{
			shared.ScriptHtmx,
			shared.ScriptSwal,
		},
	})

	return hb.Raw(layout)
}

func (c *homeController) checkAndProcess(config shared.Config) string {
	data, errorMessage := c.prepareData(config)

	if errorMessage != "" {
		return hb.Div().
			Class("alert alert-danger").
			Text(errorMessage).ToHTML()

	}

	return c.page(data).ToHTML()
}

// == PRIVATE METHODS ==========================================================

func (c *homeController) prepareData(config shared.Config) (data homeControllerData, errorMessage string) {
	return homeControllerData{
		Config: config,
	}, ""
}

func (c *homeController) page(data homeControllerData) hb.TagInterface {
	breadcrumbs := shared.Breadcrumbs(data.Config, []shared.Breadcrumb{
		{
			Name: "Home",
			URL:  shared.Url(data.Request, data.HomeURL, nil),
		},
		{
			Name: "Users",
			URL:  shared.Url(data.Request, shared.PathHome, nil),
		},
	})

	title := hb.Heading1().
		HTML("Users. Home")

	options :=
		hb.Section().
			Class("mb-3 mt-3").
			Style("background-color: #f8f9fa;").
			Child(
				hb.UL().
					Class("list-group").
					Child(hb.LI().
						Class("list-group-item").
						Child(hb.A().
							Href(shared.Url(data.Request, shared.PathUsers, nil)).
							Text("Users"))))
		// Child(hb.LI().
		// 	Class("list-group-item").
		// 	Child(hb.A().
		// 		Href(url(data.Request, pathGroups, nil)).
		// 		Text("Groups"))))

	return hb.Div().
		Class("container").
		Child(breadcrumbs).
		Child(hb.HR()).
		Child(shared.AdminHeader(data.Config)).
		Child(hb.HR()).
		Child(title).
		Child(options)
}
