package user

import (
	"net/http"
	"time"

	"github.com/ostamand/url/web/store"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const SessionCookie = "session_token"
const ExpirationDelay = time.Minute * 60 * 3

func GetFromSession(s *store.StorageService, req *http.Request) *store.UserModel {
	user := &store.UserModel{}
	if c, err := req.Cookie(SessionCookie); err == nil {
		user, _ = (*s).GetUserBySession(c.Value)
	}
	return user
}

func IsLoggedIn(storage *store.StorageService, sessionToken string) (*store.UserModel, error) {
	return (*storage).GetUserBySession(sessionToken)
}

func CreateSession(storage *store.StorageService, user *store.UserModel) (*store.SessionModel, error) {
	sessionToken := uuid.NewV4().String()
	session := store.SessionModel{
		UserID:    user.ID,
		Token:     sessionToken,
		CreatedAt: time.Now(),
		ExpiryAt:  time.Now().Add(ExpirationDelay),
	}
	err := (*storage).SaveSession(&session)
	return &session, err
}

func VerifyPassword(storage *store.StorageService, userName string, password string) (*store.UserModel, error) {
	user, err := (*storage).GetByUsername(userName)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
