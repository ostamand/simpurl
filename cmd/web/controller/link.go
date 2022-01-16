package controller

import (
	"net/http"

	"github.com/ostamand/simpurl/cmd/web/notify"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/user"
)

type LinkController struct {
	Storage *store.StorageService
	User    *user.UserHelper
}

func (c *LinkController) List(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u, ok := c.User.HasAccess(w, req, "/signin")
		if !ok {
			return
		}
		links, _ := c.Storage.Link.GetAll(u.ID)
		viewData := CreateViewData(req, u)
		data := struct {
			*ViewData
			Links *[]store.LinkModel
		}{
			viewData,
			links,
		}
		ShowPage(w, data, "link/list.page.gohtml")
	}
}

func (c LinkController) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u, ok := c.User.HasAccess(w, req, "/login")
		if !ok {
			return
		}
		data := CreateViewData(req, u)
		ShowPage(w, data, "link/create.page.html")

	case http.MethodPost:
		u := c.User.GetFromSession(req)
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		l := &store.LinkModel{
			UserID:      u.ID,
			Symbol:      req.FormValue("symbol"),
			URL:         req.FormValue("url"),
			Description: req.FormValue("description"),
			Note:        req.FormValue("note"),
		}

		// TODO check if URL already exists
		// TODO check if symbol already associated
		c.Storage.Link.Save(l)

		url := notify.AddNotificationToURL("/link/create", notify.NotifyLinkCreated)
		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}
