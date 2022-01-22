package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/user"
)

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserController struct {
	Storage *store.StorageService
	User    *user.UserHelper
}

func (c *UserController) Signin(w http.ResponseWriter, req *http.Request) {
	AllowOrigins(&w)
	switch req.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		request := SigninRequest{}
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()

		_ = json.Unmarshal(body, &request)

		username:= request.Username
		password:= request.Password

		// check password
		u, err := c.User.VerifyPassword(username, password)
		if err != nil || (c.User.AdminOnly && !u.Admin) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// create session token
		session, err := c.User.CreateSession(w, u.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := struct {
			Token string `json:"token"`
		}{
			Token: session.Token,
		}
		json.NewEncoder(w).Encode(data)
	}
}
