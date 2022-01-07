package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ostamand/simpurl/internal/config"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/store/mysql"
	"github.com/ostamand/simpurl/internal/user"
)

var userCtrl *UserController

func init() {
	wd, _ := os.Getwd()
	configPath, _ := config.FindIn(wd, os.Getenv("CONFIG_FILE"))
	params := config.Get(configPath)
	storage := mysql.InitializeSQL(&params.Db)

	u := &user.UserHelper{AdminOnly: false, Storage: storage}
	userCtrl = &UserController{Storage: storage, User: u}
}

func sendSigninRequest(username string, password string) (*http.Response, SigninResponse) {
	dataRequest := SigninRequest{
		Username: username,
		Password: password,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(dataRequest)

	req, _ := http.NewRequest(http.MethodPost, "/api/signin", b)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	userCtrl.Signin(w, req)

	resp := w.Result()
	dataResponse := SigninResponse{}
	json.NewDecoder(w.Body).Decode(&dataResponse)

	return resp, dataResponse
}

func TestSigninUserDoesNotExists(t *testing.T) {
	// user does not exists
	resp, data := sendSigninRequest("user", "password")
	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("User does not exists. Should get a 401")
	}
	if data.Token != "" {
		t.Errorf("Expecting no token but got %s", data.Token)
	}
}

func TestAdmin(t *testing.T) {
	userCtrl.User.AdminOnly = true

	u := &store.UserModel{
		Username: "user",
		Password: "password",
		Admin:    false,
	}
	userCtrl.Storage.User.Save(u)
	resp, data := sendSigninRequest(u.Username, u.Password)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Expection status 401 since users is not admin")
	}
	if data.Token != "" {
		t.Errorf("Expecting no token but got %s", data.Token)
	}

	// cleanup
	userCtrl.Storage.User.DeleteFromUsername(u.Username)
	userCtrl.User.AdminOnly = false
}

func TestSigninUserExists(t *testing.T) {
	u := &store.UserModel{
		Username: "user",
		Password: "password",
		Admin:    false,
	}
	userCtrl.Storage.User.Save(u)
	resp, data := sendSigninRequest(u.Username, u.Password)
	if resp.StatusCode != http.StatusOK {
		t.Error("Expection status 200 since users exists")
	}
	if data.Token == "" {
		t.Errorf("Expeciting a session token to be provided but got %s", data.Token)
	}

	userFromSession, err := userCtrl.Storage.User.GetBySession(data.Token)
	if err != nil {
		t.Error(err)
	}

	if userFromSession.Username != u.Username {
		t.Errorf("Expecting to get username %s from session but got %s", u.Username, userFromSession.Username)
	}

	// cleanup
	err = userCtrl.Storage.Session.DeleteFromToken(data.Token)
	if err != nil {
		t.Error(err)
	}
	err = userCtrl.Storage.User.Delete(userFromSession.ID)
	if err != nil {
		t.Error(err)
	}
}
