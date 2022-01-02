package helper

import (
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ostamand/url/web/config"
	"github.com/ostamand/url/web/store"
	"github.com/ostamand/url/web/store/mysql"
	"github.com/stretchr/testify/assert"
)

type SessionMock struct {
	Token string
}

func (s *SessionMock) Save(http.ResponseWriter) (string, time.Time) {
	sessionToken, expires := GenerateToken()
	s.Token = sessionToken
	return sessionToken, expires
}

func (s *SessionMock) Get(*http.Request) (string, error) {
	var err error = nil
	if s.Token == "" {
		err = errors.New("No session saved")
	}
	return s.Token, err
}

func (s *SessionMock) Clear(http.ResponseWriter) {
	s.Token = ""
}

var storage *store.StorageService

func init() {
	wd, _ := os.Getwd()
	configPath, _ := config.FindIn(wd, os.Getenv("CONFIG_FILE"))
	params := config.Get(configPath)
	storage = mysql.InitializeSQL(&params.Db)
}

func TestAdminNoAccess(t *testing.T) {
	h := UserHelper{
		AdminOnly: true,
		Storage:   storage,
		Session:   &SessionMock{},
	}

	u := &store.UserModel{
		Admin:    false,
		Username: "user",
		Password: "test",
	}

	// save user
	h.Storage.User.Save(u)

	// get from db, we need the id
	u, _ = h.Storage.User.GetByUsername(u.Username)

	// login user
	s, _ := h.CreateSession(nil, u.ID)

	// check user from session
	uFromSession := h.GetFromSession(nil)
	assert.Equal(t, u.ID, uFromSession.ID)

	req, _ := http.NewRequest("GET", "/link", nil)
	_, ok := h.HasAccess(nil, req, "signin")
	assert.False(t, ok)

	// cleanup
	h.Storage.User.Delete(u.ID)
	h.Storage.Session.DeleteFromToken(s.Token)

}
