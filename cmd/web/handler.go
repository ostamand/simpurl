package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	ctrl "github.com/ostamand/url/cmd/web/controller"
	"github.com/ostamand/url/cmd/web/helper"
	"github.com/ostamand/url/cmd/web/store"
)

type Handler struct {
	Storage *store.StorageService
	User    *helper.UserHelper
}

func (h *Handler) home(w http.ResponseWriter, req *http.Request) {
	u, ok := h.User.HasAccess(w, req, "/signin")
	if ok {
		data := ctrl.CreateViewData(req, u)
		ctrl.ShowPage(w, data, "home.page.html")
	}
}

func (h *Handler) redirect(w http.ResponseWriter, req *http.Request) {
	url := req.URL.String()

	if url == "/" {
		h.home(w, req)
		return
	}

	if _, ok := h.User.HasAccess(w, req, "/signin"); !ok {
		return
	}

	splits := strings.Split(url, "/")[1:]
	if len(splits) > 1 {
		http.Error(w, "Bad request, expecting: /shorturl", http.StatusBadRequest)
		return
	}

	symbol := splits[0]
	if l, err := h.Storage.Link.FindBySymbol(symbol); err != nil {
		log.Printf("error during redirect: %s", err)
		text := fmt.Sprintf("Short URL not found: %s", symbol)
		http.Error(w, text, http.StatusBadRequest)
		return
	} else {
		http.Redirect(w, req, l.URL, http.StatusSeeOther)
	}
}
