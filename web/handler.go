package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ostamand/url/web/notify"
	"github.com/ostamand/url/web/store"
	"github.com/ostamand/url/web/user"
)

type Handler struct {
	storage store.StorageService
}

func showPage(w io.Writer, data interface{}, pages ...string) {
	for i, p := range pages {
		pages[i] = "ui/html/" + p
	}
	pages = append(pages,
		"ui/html/base.layout.html",
		"ui/html/logged.partial.html",
		"ui/html/notify.partial.html",
	)
	tmpl := template.Must(template.ParseFiles(pages...))
	tmpl.Execute(w, data)
}

func (h Handler) signup(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		showPage(w, nil, "signup.page.html")
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
		Name:   SessionCookie,
		MaxAge: -1,
	})
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func (h Handler) signin(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		data := CreateViewModel(req, nil)
		showPage(w, data, "signin.page.html")
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
			Name:    SessionCookie,
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
	data := CreateViewModel(req, u)
	showPage(w, data, "home.page.html")
}

func (h Handler) redirect(w http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	if url == "/" {
		h.home(w, req)
		return
	}
	splits := strings.Split(req.URL.String(), "/")[1:]
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

func (h Handler) links(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u := user.GetFromSession(&h.storage, req)
		if u.Authenticated() {
			data := CreateViewModel(req, u)
			showPage(w, data, "links.page.html")
			return
		} else {
			http.Redirect(w, req, "/signin", http.StatusSeeOther)
		}
	case http.MethodPost:
		u := user.GetFromSession(&h.storage, req)
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
		h.storage.SaveLink(l)

		url := notify.AddNotificationToURL("/links", notify.NotifyLinkCreated)
		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}
