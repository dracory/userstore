package admin

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/userstore/admin/shared"
	"github.com/samber/lo"
	"github.com/spf13/cast"
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
	breadcrumbs := shared.Breadcrumbs(data.Config, []shared.Breadcrumb{})

	title := hb.Heading1().
		Text("User Management Dashboard")

	return hb.Div().
		Class("container").
		Child(breadcrumbs).
		Child(hb.HR()).
		Child(shared.AdminHeader(data.Config)).
		Child(hb.HR()).
		Child(title).
		Children(tiles(data))
}

// == PUBLIC METHODS ===========================================================

func tiles(data homeControllerData) []hb.TagInterface {
	usersCount, errUsersCount := data.Store.UserCount(data.Request.Context(), userstore.NewUserQuery())

	if errUsersCount != nil {
		usersCount = 0
	}

	tiles := []struct {
		Count      string
		Title      string
		Background string
		Icon       string
		URL        string
	}{

		{
			Count:      cast.ToString(usersCount),
			Title:      "Total Users",
			Background: "bg-success",
			Icon:       "bi-users",
			URL:        shared.Url(data.Request, shared.PathUsers, nil),
		},
		// {
		// 	Count:      cast.ToString(pagesCount),
		// 	Title:      "Total Pages",
		// 	Background: "bg-info",
		// 	Icon:       "bi-journals",
		// 	URL:        shared.URLR(r, shared.PathPagesPageManager, nil),
		// },
		// {
		// 	Count:      cast.ToString(templatesCount),
		// 	Title:      "Total Templates",
		// 	Background: "bg-warning",
		// 	Icon:       "bi-file-earmark-text-fill",
		// 	URL:        shared.URLR(r, shared.PathTemplatesTemplateManager, nil),
		// },
		// {
		// 	Count:      cast.ToString(blocksCount),
		// 	Title:      "Total Blocks",
		// 	Background: "bg-primary",
		// 	Icon:       "bi-grid-3x3-gap-fill",
		// 	URL:        shared.URLR(r, shared.PathBlocksBlockManager, nil),
		// },
	}

	cards := lo.Map(tiles, func(tile struct {
		Count      string
		Title      string
		Background string
		Icon       string
		URL        string
	}, index int) hb.TagInterface {
		card := hb.Div().
			Class("card").
			Class("bg-transparent border round-10 shadow-lg h-100").
			// OnMouseOver(`this.style.setProperty('background-color', 'beige', 'important');this.style.setProperty('scale', 1.1);this.style.setProperty('border', '4px solid moccasin', 'important');`).
			// OnMouseOut(`this.style.setProperty('background-color', 'transparent', 'important');this.style.setProperty('scale', 1);this.style.setProperty('border', '4px solid transparent', 'important');`).
			Child(hb.Div().
				Class("card-body").
				Class(tile.Background).
				Style("--bs-bg-opacity:0.3;").
				Child(hb.Div().Class("row").
					Child(hb.Div().Class("col-8").
						Child(hb.Div().
							Style("margin-top:-4px;margin-right:8px;font-size:32px;").
							Text(tile.Count)).
						Child(hb.NewDiv().
							Style("margin-top:-4px;margin-right:8px;font-size:16px;").
							Text(tile.Title)),
					).
					Child(hb.Div().Class("col-4").
						Child(hb.I().
							Class("bi float-end").
							Class(tile.Icon).
							Style(`color:silver;opacity:0.6;`).
							Style("margin-top:-4px;margin-right:8px;font-size:48px;")),
					),
				)).
			Child(hb.Div().
				Class("card-footer text-center").
				Class(tile.Background).
				Style("--bs-bg-opacity:0.5;").
				Child(hb.A().
					Class("text-white").
					Href(tile.URL).
					Text("More info").
					Child(hb.I().Class("bi bi-arrow-right-circle-fill ms-3").Style("margin-top:-4px;margin-right:8px;font-size:16px;")),
				))
		return hb.Div().Class("col-xs-12 col-sm-6 col-md-3").Child(card)
	})

	return cards
}
