package helper

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const SessionCookie = "session_token"
const ExpirationDelay = time.Minute * 60 * 3

type SessionClient interface {
	Get(req *http.Request) (string, error)
	Save(w http.ResponseWriter) (string, time.Time)
	Clear(w http.ResponseWriter)
}

type SessionHTTP struct{}

func GenerateToken() (token string, expires time.Time) {
	token = uuid.NewV4().String()
	expires = time.Now().Add(ExpirationDelay)
	return
}

func (s SessionHTTP) Get(req *http.Request) (string, error) {
	c, err := req.Cookie(SessionCookie)
	if err == nil {
		return c.Value, err
	}
	return "", err
}

func (s SessionHTTP) Save(w http.ResponseWriter) (string, time.Time) {
	sessionToken, expires := GenerateToken()
	http.SetCookie(w, &http.Cookie{
		Name:    SessionCookie,
		Value:   sessionToken,
		Expires: expires,
	})
	return sessionToken, expires
}

func (s SessionHTTP) Clear(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   SessionCookie,
		MaxAge: -1,
	})
}
