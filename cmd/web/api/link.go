package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ostamand/simpurl/cmd/web/helper"
	"github.com/ostamand/simpurl/internal/store"
)

type LinkController struct {
	Storage *store.StorageService
	User    *helper.UserHelper
}

type CreateRequest struct {
	Token       string `json:"token"`
	Symbol      string `json:"symbol"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Note        string `json:"note"`
}

func (c *LinkController) Create(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	request := CreateRequest{}
	_ = json.Unmarshal(body, &request)

	u, err := c.Storage.User.GetBySession(request.Token)
	if err != nil || (c.User.AdminOnly && !u.Admin) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	l := &store.LinkModel{
		UserID:      u.ID,
		Symbol:      request.Symbol,
		URL:         request.URL,
		Description: request.Description,
		Note:        request.Note,
	}
	if err = c.Storage.Link.Save(l); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
