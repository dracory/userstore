package shared

import (
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

type Breadcrumb struct {
	Name string
	URL  string
}

func Breadcrumbs(config Config, pageBreadcrumbs []Breadcrumb) hb.TagInterface {
	adminHomeURL := config.HomeURL

	adminHomeBreadcrumb := lo.
		If(adminHomeURL != "", Breadcrumb{
			Name: "Home",
			URL:  adminHomeURL,
		}).
		Else(Breadcrumb{})

	breadcrumbItems := []Breadcrumb{
		adminHomeBreadcrumb,
		{
			Name: "Users",
			URL:  Url(config.Request, PathHome, nil),
		},
	}

	breadcrumbItems = append(breadcrumbItems, pageBreadcrumbs...)

	breadcrumbs := breadcrumbsUI(breadcrumbItems)

	return hb.Div().
		Child(breadcrumbs)
}

func breadcrumbsUI(breadcrumbs []Breadcrumb) hb.TagInterface {

	ol := hb.OL().
		Class("breadcrumb").
		Style("margin-bottom: 0px;")

	for _, breadcrumb := range breadcrumbs {
		link := hb.Hyperlink().
			HTML(breadcrumb.Name).
			Href(breadcrumb.URL)

		li := hb.LI().
			Class("breadcrumb-item").
			Child(link)

		ol.AddChild(li)
	}

	nav := hb.Nav().
		Class("d-inline-block").
		Attr("aria-label", "breadcrumb").
		Child(ol)

	return nav
}
