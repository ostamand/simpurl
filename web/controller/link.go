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

func (c LinkController) List(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u := user.GetFromSession(&c.Storage, req)
		if !u.Authenticated() {
			url := notify.AddNotificationToURL("/signin", notify.NotifyNotSignedIn)
			http.Redirect(w, req, url, http.StatusSeeOther)
			return
		}
		links, _ := c.Storage.GetAllLinks(u)
		viewData := CreateViewData(req, u)
		data := struct {
			*ViewData
			Links *[]store.LinkModel
		}{
			viewData,
			links,
		}
		ShowPage(w, data, "link/list.page.html")
	}
}

func (c LinkController) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u := user.GetFromSession(&c.Storage, req)
		if !u.Authenticated() {
			url := notify.AddNotificationToURL("/signin", notify.NotifyNotSignedIn)
			http.Redirect(w, req, url, http.StatusSeeOther)
			return
		}
		data := CreateViewData(req, u)
		ShowPage(w, data, "link/create.page.html")

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

		url := notify.AddNotificationToURL("/link/create", notify.NotifyLinkCreated)
		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}
