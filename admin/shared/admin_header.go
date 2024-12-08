package shared

import (
	"context"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/userstore"
	"github.com/spf13/cast"
)

func AdminHeader(config Config) hb.TagInterface {
	linkHome := hb.NewHyperlink().
		HTML("Dashboard").
		Href(Url(config.Request, PathHome, nil)).
		Class("nav-link")
	linkUsers := hb.Hyperlink().
		HTML("Users").
		Href(Url(config.Request, PathUsers, nil)).
		Class("nav-link")
	// linkTasks := hb.Hyperlink().
	// 	HTML("Tasks").
	// 	Href(url(r, pathTaskManager, nil)).
	// 	Class("nav-link")

	userCount, err := config.Store.UserCount(context.Background(), userstore.NewUserQuery())

	if err != nil {
		config.Logger.Error(err.Error())
		userCount = -1
	}

	// taskCount, err := store.TaskCount(statsstore.TaskQuery())

	// if err != nil {
	// 	logger.Error(err.Error())
	// 	taskCount = -1
	// }

	ulNav := hb.NewUL().Class("nav  nav-pills justify-content-center")
	ulNav.AddChild(hb.NewLI().Class("nav-item").Child(linkHome))

	ulNav.Child(hb.LI().
		Class("nav-item").
		Child(linkUsers.
			Child(hb.Span().
				Class("badge bg-secondary ms-2").
				HTML(cast.ToString(userCount)))))

	// ulNav.Child(hb.LI().
	// 	Class("nav-item").
	// 	Child(linkTasks.
	// 		Child(hb.Span().
	// 			Class("badge bg-secondary ms-2").
	// 			HTML(cast.ToString(taskCount)))))

	divCard := hb.NewDiv().Class("card card-default mt-3 mb-3")
	divCardBody := hb.NewDiv().Class("card-body").Style("padding: 2px;")
	return divCard.AddChild(divCardBody.AddChild(ulNav))
}

// func redirect(w http.ResponseWriter, r *http.Request, url string) string {
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// 	http.Redirect(w, r, url, http.StatusSeeOther)
// 	return ""
// }
