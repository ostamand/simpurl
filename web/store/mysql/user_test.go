package mysql

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ostamand/url/web/config"
	"github.com/ostamand/url/web/store"
	"github.com/stretchr/testify/assert"
)

var storage *store.StorageService

func init() {
	wd, _ := os.Getwd()
	fullPath := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(wd))), os.Getenv("CONFIG_FILE"))

	params := config.Get(fullPath)
	storage = InitializeSQL(&params.Db)
}

func TestSave(t *testing.T) {
	table := []store.UserModel{
		{
			Username:  "username_1",
			Password:  "password_1",
			Admin:     false,
			CreatedAt: time.Now(),
		},
		{
			Username:  "username_2",
			Password:  "password_2",
			Admin:     true,
			CreatedAt: time.Now(),
		},
	}
	for _, u := range table {
		t.Run(u.Username, func(*testing.T) {
			storage.User.Save(&u)
			uActual, err := storage.User.GetByUsername(u.Username)
			assert.NoError(t, err)
			if assert.NoError(t, err) {
				assert.Equal(t, u.Username, uActual.Username)
				assert.Equal(t, u.Admin, uActual.Admin)
			}
			assert.NoError(t, storage.User.Delete(uActual.ID))
			_, err = storage.User.GetByUsername(u.Username)
			assert.Error(t, err)
		})
	}
}
