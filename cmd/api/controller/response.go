package controller

import "github.com/ostamand/simpurl/internal/store"

type ListResponse struct {
	Count int `json:"count"`
	Links []store.LinkModel `json:"links"`
}

type RedirectResponse struct {
	URL string `json:"url"`
}