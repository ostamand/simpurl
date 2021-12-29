package controller

import (
	"net/http"

	"github.com/ostamand/url/web/notify"
	"github.com/ostamand/url/web/store"
	"github.com/ostamand/url/web/user"
)

type LinkController struct {
	Storage store.StorageService
}

func (c LinkController) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u := user.GetFromSession(&c.Storage, req)
		if u.Authenticated() {
			data := CreateViewData(req, u)
			ShowPage(w, data, "links.page.html")
			return
		} else {
			url := notify.AddNotificationToURL("/signin", notify.NotifyNotSignedIn)
			http.Redirect(w, req, url, http.StatusSeeOther)
		}
	case http.MethodPost:
		u := user.GetFromSession(&c.Storage, req)
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		l := &store.LinkModel{
			UserID:      u.ID,
			Symbol:      req.FormValue("symbol"),
			URL:         req.FormValue("url"),
			Description: req.FormValue("description"),
		}

		// TODO check if URL already exists
		// TODO check if symbol already associated
		c.Storage.SaveLink(l)

		url := notify.AddNotificationToURL("/links/create", notify.NotifyLinkCreated)
		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}
