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

type ListRequest struct {
	Token       string 	`json:"token"`
	Limit		int 	`json:"limit"`
}

type ListResponse struct {
	Count int `json:"count"`
	Links []store.LinkModel `json:"links"`
}

func allowPostRequest(w *http.ResponseWriter, req *http.Request) bool {
	controller.AllowOrigins(w)

	if req.Method == http.MethodOptions {
		(*w).WriteHeader(http.StatusOK)
		return true
	}
	if req.Method != http.MethodPost {
		(*w).WriteHeader(http.StatusNotFound)
		return false
	}
	return true
}

func (c *LinkController) getUser(token string) (*store.UserModel, bool) {
	u, err := c.Storage.User.GetBySession(token)
	if err != nil || (c.User.AdminOnly && !u.Admin) {
		return u, false
	}
	return u, true
}

// TODO support filters
func (c *LinkController) List(w http.ResponseWriter, req *http.Request) {
	if ok := allowPostRequest(&w, req); !ok {
		return 
	}

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	request := ListRequest{}
	_ = json.Unmarshal(body, &request)

	u, ok := c.getUser(request.Token)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}

	// TODO implement limit on GetAll
	links, err := c.Storage.Link.GetAll(u.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO setup count
	response := ListResponse {
		Count: 0,
		Links: *links,
	}

	json.NewEncoder(w).Encode(response)
}

func (c *LinkController) Create(w http.ResponseWriter, req *http.Request) {
	if ok := allowPostRequest(&w, req); !ok {
		return 
	}

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	request := CreateRequest{}
	_ = json.Unmarshal(body, &request)

	u, ok := c.getUser(request.Token)
	if !ok {
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
	if err := c.Storage.Link.Save(l); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
