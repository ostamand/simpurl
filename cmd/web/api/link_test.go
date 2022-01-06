package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/ostamand/simpurl/cmd/web/helper"
	"github.com/ostamand/simpurl/internal/config"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/store/mysql"
)

var linkCtrl *LinkController

func init() {
	wd, _ := os.Getwd()
	configPath, _ := config.FindIn(wd, os.Getenv("CONFIG_FILE"))
	params := config.Get(configPath)
	storage := mysql.InitializeSQL(&params.Db)

	u := &helper.UserHelper{AdminOnly: false, Storage: storage}
	linkCtrl = &LinkController{Storage: storage, User: u}
}

func sendRequest(sessionToken string, l *store.LinkModel) *http.Response {
	dataRequest := CreateRequest{
		Token:       sessionToken,
		Symbol:      l.Symbol,
		URL:         l.URL,
		Description: l.Description,
		Note:        l.Note,
	}
	b, _ := json.Marshal(dataRequest)

	req, _ := http.NewRequest(http.MethodPost, "/api/link/create", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	linkCtrl.Create(w, req)

	return w.Result()
}

func createUserAndSession() (*store.UserModel, *store.SessionModel) {
	u := &store.UserModel{
		Username: "user",
		Password: "password",
		Admin:    false,
	}
	userCtrl.Storage.User.Save(u)
	u, _ = userCtrl.Storage.User.GetByUsername(u.Username)

	token, expires := helper.GenerateToken()
	session := &store.SessionModel{
		UserID:    u.ID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiryAt:  expires,
	}
	userCtrl.Storage.Session.Save(session)
	return u, session
}

func cleanupUserAndSession(username string, token string) {
	userCtrl.Storage.User.DeleteFromUsername(username)
	userCtrl.Storage.Session.DeleteFromToken(token)
}

func TestCreateNoToken(t *testing.T) {
	// user and session exists but token not provided
	u, session := createUserAndSession()
	defer cleanupUserAndSession(u.Username, session.Token)

	l := &store.LinkModel{}
	resp := sendRequest("", l)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("No session token provided, should get a 401")
	}
}

func TestCreateNoSession(t *testing.T) {
	u, session := createUserAndSession()
	userCtrl.Storage.Session.DeleteFromToken(session.Token)
	defer func() { userCtrl.Storage.User.DeleteFromUsername(u.Username) }()

	l := &store.LinkModel{}
	resp := sendRequest(session.Token, l)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("No session exists for that user, should get a 401")
	}
}

func TestCreateWithSession(t *testing.T) {
	u, session := createUserAndSession()
	defer cleanupUserAndSession(u.Username, session.Token)

	l := &store.LinkModel{
		UserID:      u.ID,
		URL:         "https://www.google.com/robots.txt",
		Description: "Robots on Google",
		Note:        "Run!",
	}
	defer func() { userCtrl.Storage.Link.DeleteByURL(l.URL) }()
	resp := sendRequest(session.Token, l)

	if resp.StatusCode != http.StatusOK {
		t.Error("User with active session, should get 200")
	}
}
