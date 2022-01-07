package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ostamand/simpurl/internal/session"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/user"
)

type UserController struct {
	Storage *store.StorageService
	User    *user.UserHelper
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Token string `json:"token"`
}

func (c *UserController) Signin(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	request := SigninRequest{}
	_ = json.Unmarshal(body, &request)

	u, err := c.User.VerifyPassword(request.Username, request.Password)
	if err != nil || (c.User.AdminOnly && !u.Admin) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create new session
	token, expires := session.GenerateToken()
	session := store.SessionModel{
		UserID:    u.ID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiryAt:  expires,
	}
	err = c.Storage.Session.Save(&session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// prepare response
	data := SigninResponse{Token: token}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
