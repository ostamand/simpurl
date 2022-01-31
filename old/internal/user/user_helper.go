package user

import (
	"net/http"
	"time"

	"github.com/ostamand/simpurl/internal/session"
	"github.com/ostamand/simpurl/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type UserHelper struct {
	AdminOnly bool
	Storage   *store.StorageService
}

func (h *UserHelper) CreateSession(w http.ResponseWriter, userID int) (*store.SessionModel, error) {
	sessionToken, expires := session.GenerateToken()
	session := store.SessionModel{
		UserID:    userID,
		Token:     sessionToken,
		CreatedAt: time.Now(),
		ExpiryAt:  expires,
	}
	err := h.Storage.Session.Save(&session)
	return &session, err
}

func (h *UserHelper) VerifyPassword(username string, password string) (*store.UserModel, error) {
	user, err := h.Storage.User.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return user, err
}
