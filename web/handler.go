package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ostamand/url/web/store"
)

type CreateRequest struct {
	Symbol      string `json:"symbol"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type Handler struct {
	storage store.StorageService
}

func showPage(page string, w io.Writer, data interface{}) {
	templates := []string{page, "ui/html/base.layout.html"}
	tmpl := template.Must(template.ParseFiles(templates...))
	tmpl.Execute(w, data)
}

func (h Handler) signup(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		showPage("ui/html/signup.page.html", w, nil)
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

func (h Handler) signin(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// TODO shows signout if already logged in
		showPage("ui/html/signin.page.html", w, nil)
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		username := req.FormValue("username")
		password := req.FormValue("password")

		// check password
		user, err := store.VerifyPassword(&h.storage, username, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			// TODO redirect to signin
			return
		}

		// create session token
		session, err := store.CreateSession(&h.storage, user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO redirect to signin
			return
		}

		// set cookie based on session
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   session.Token,
			Expires: session.ExpiryAt,
		})

		// go back to home for now
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

func (h Handler) redirect(w http.ResponseWriter, req *http.Request) {
	if req.URL.String() == "/" {
		fmt.Fprintf(w, "Welcome to Short URL")
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

func (h Handler) create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Printf("GET")
	case "POST":
		contentType := req.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			var r CreateRequest
			err := json.NewDecoder(req.Body).Decode(&r)
			if err != nil {
				http.Error(w, "Bad request json", http.StatusBadRequest)
				return
			}

			l := &store.LinkModel{
				Symbol:      r.Symbol,
				URL:         r.URL,
				Description: r.Description,
			}

			h.storage.SaveLink(l)

			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Short URL created for: %s -> %s (%s)", r.Symbol, r.URL, r.Description)
		}
	}
}
