package user

import (
	"net/http"
	"os"
	"testing"

	"github.com/ostamand/simpurl/internal/config"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/store/mysql"
	"github.com/ostamand/simpurl/test"
	"github.com/stretchr/testify/assert"
)

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
		Session:   &test.SessionMock{},
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
