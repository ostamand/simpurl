package helper

import (
	"net/http"
	"time"

	"github.com/ostamand/simpurl/cmd/web/notify"
	"github.com/ostamand/simpurl/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type UserHelper struct {
	AdminOnly bool
	Storage   *store.StorageService
	Session   SessionClient
}

func (h *UserHelper) HasAccess(w http.ResponseWriter, req *http.Request, redirect string) (*store.UserModel, bool) {
	u := h.GetFromSession(req)
	adminOnly := h.AdminOnly && !u.Admin
	if !u.Authenticated() || adminOnly {
		var url string
		if adminOnly {
			url = notify.AddNotificationToURL(redirect, notify.NotifyInDev)
		} else {
			url = notify.AddNotificationToURL(redirect, notify.NotifyNotSignedIn)
		}
		if w != nil {
			http.Redirect(w, req, url, http.StatusSeeOther)
		}
		return nil, false
	}
	return u, true
}

func (h *UserHelper) GetFromSession(req *http.Request) *store.UserModel {
	user := &store.UserModel{}
	if c, err := h.Session.Get(req); err == nil {
		user, _ = h.Storage.User.GetBySession(c)
	}
	return user
}

func (h *UserHelper) CreateSession(w http.ResponseWriter, userID int) (*store.SessionModel, error) {
	sessionToken, expires := h.Session.Save(w)
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
