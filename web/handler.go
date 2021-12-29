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
