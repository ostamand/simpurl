package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	ctrl "github.com/ostamand/url/web/controller"
	"github.com/ostamand/url/web/notify"
	"github.com/ostamand/url/web/store"
	"github.com/ostamand/url/web/user"
)

type Handler struct {
	storage store.StorageService
}

func (h Handler) signup(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ctrl.ShowPage(w, nil, "signup.page.html")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := h.storage.SaveUser(req.FormValue("username"), req.FormValue("password"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

func (h Handler) signout(w http.ResponseWriter, req *http.Request) {
	// delete cookie
	http.SetCookie(w, &http.Cookie{
		Name:   ctrl.SessionCookie,
		MaxAge: -1,
	})
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func (h Handler) signin(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		data := ctrl.CreateViewData(req, nil)
		ctrl.ShowPage(w, data, "signin.page.html")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		username := req.FormValue("username")
		password := req.FormValue("password")

		// check password
		u, err := user.VerifyPassword(&h.storage, username, password)
		if err != nil {
			url := notify.AddNotificationToURL("/signin", notify.NotifyWrongPassword)
			http.Redirect(w, req, url, http.StatusSeeOther)
			return
		}

		// create session token
		session, err := user.CreateSession(&h.storage, u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO redirect to signin
			return
		}

		// set cookie based on session
		http.SetCookie(w, &http.Cookie{
			Name:    ctrl.SessionCookie,
			Value:   session.Token,
			Expires: session.ExpiryAt,
		})

		// go back to home for now
		http.Redirect(
			w,
			req,
			notify.AddNotificationToURL("/home", notify.NotifySignedIn),
			http.StatusSeeOther,
		)
	}
}

func (h Handler) home(w http.ResponseWriter, req *http.Request) {
	u := user.GetFromSession(&h.storage, req)
	data := ctrl.CreateViewData(req, u)
	ctrl.ShowPage(w, data, "home.page.html")
}

func (h Handler) redirect(w http.ResponseWriter, req *http.Request) {
	url := req.URL.String()

	if url == "/" {
		h.home(w, req)
		return
	}

	u := user.GetFromSession(&h.storage, req)
	if !u.Authenticated() {
		url := notify.AddNotificationToURL("/signin", notify.NotifyNotSignedIn)
		http.Redirect(w, req, url, http.StatusSeeOther)
		return
	}

	splits := strings.Split(url, "/")[1:]
	if len(splits) > 1 {
		http.Error(w, "Bad request, expecting: /shorturl", http.StatusBadRequest)
		return
	}

	symbol := splits[0]
	if l, err := h.storage.FindBySymbol(symbol); err != nil {
		log.Printf("error during redirect: %s", err)
		text := fmt.Sprintf("Short URL not found: %s", symbol)
		http.Error(w, text, http.StatusBadRequest)
		return
	} else {
		http.Redirect(w, req, l.URL, http.StatusSeeOther)
	}
}
