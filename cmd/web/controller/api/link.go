package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ostamand/simpurl/cmd/web/controller"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/user"
)

type LinkController struct {
	Storage *store.StorageService
	User    *user.UserHelper
}

type CreateRequest struct {
	Token       string `json:"token"`
	Symbol      string `json:"symbol"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Note        string `json:"note"`
}

func (c *LinkController) Create(w http.ResponseWriter, req *http.Request) {
	controller.AllowOrigins(&w) // don't use cookies anyway

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
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
	w.WriteHeader(http.StatusOK)
}
