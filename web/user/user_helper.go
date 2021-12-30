package user

import (
	"net/http"
	"time"

	"github.com/ostamand/url/web/notify"
	"github.com/ostamand/url/web/store"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHelper struct {
	AdminOnly bool
	store.StorageService
}

const SessionCookie = "session_token"
const ExpirationDelay = time.Minute * 60 * 3

func (h UserHelper) HasAccess(w http.ResponseWriter, req *http.Request, redirect string) bool {
	u := h.GetFromSession(req)
	if !u.Authenticated() {
		url := notify.AddNotificationToURL(redirect, notify.NotifyNotSignedIn)
		http.Redirect(w, req, url, http.StatusSeeOther)
		return false
	}
	return true
}

func (h UserHelper) GetFromSession(req *http.Request) *store.UserModel {
	user := &store.UserModel{}
	if c, err := req.Cookie(SessionCookie); err == nil {
		user, _ = h.GetUserBySession(c.Value)
	}
	return user
}

func (h UserHelper) IsLoggedIn(sessionToken string) (*store.UserModel, error) {
	return h.GetUserBySession(sessionToken)
}

func (h UserHelper) CreateSession(user *store.UserModel) (*store.SessionModel, error) {
	sessionToken := uuid.NewV4().String()
	session := store.SessionModel{
		UserID:    user.ID,
		Token:     sessionToken,
		CreatedAt: time.Now(),
		ExpiryAt:  time.Now().Add(ExpirationDelay),
	}
	err := h.SaveSession(&session)
	return &session, err
}

func (h UserHelper) VerifyPassword(username string, password string) (*store.UserModel, error) {
	user, err := h.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
