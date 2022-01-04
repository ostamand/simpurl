package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ostamand/url/web/helper"
	"github.com/ostamand/url/web/store"
)

type UserController struct {
	Storage *store.StorageService
	User    *helper.UserHelper
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
		return
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	request := SigninRequest{}
	_ = json.Unmarshal(body, &request)

	u, err := c.User.VerifyPassword(request.Username, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create new session
	token, expires := helper.GenerateToken()
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
