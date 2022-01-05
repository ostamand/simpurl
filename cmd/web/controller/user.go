package controller

import (
	"net/http"

	"github.com/ostamand/url/cmd/web/helper"
	"github.com/ostamand/url/cmd/web/notify"
	"github.com/ostamand/url/cmd/web/store"
)

type UserController struct {
	Storage *store.StorageService
	User    *helper.UserHelper
}

func (c *UserController) Signup(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ShowPage(w, nil, "signup.page.html")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		u := store.UserModel{
			Username: req.FormValue("username"),
			Password: req.FormValue("password"),
		}
		err := c.Storage.User.Save(&u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

func (c *UserController) Signout(w http.ResponseWriter, req *http.Request) {
	c.User.Session.Clear(w)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func (c *UserController) Signin(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		data := CreateViewData(req, nil)
		ShowPage(w, data, "signin.page.html")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		username := req.FormValue("username")
		password := req.FormValue("password")

		// check password
		u, err := c.User.VerifyPassword(username, password)
		if err != nil {
			url := notify.AddNotificationToURL("/signin", notify.NotifyWrongPassword)
			http.Redirect(w, req, url, http.StatusSeeOther)
			return
		}

		// create session token
		_, err = c.User.CreateSession(w, u.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO redirect to signin
			return
		}

		// go back to home for now
		http.Redirect(
			w,
			req,
			notify.AddNotificationToURL("/home", notify.NotifySignedIn),
			http.StatusSeeOther,
		)
	}
}
