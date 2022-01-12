package session

import (
	"net/http"
	"time"

	"github.com/ostamand/simpurl/internal/session"
)

const SessionCookie = "session_token"

type SessionHTTP struct{}

func (s *SessionHTTP) Get(req *http.Request) (string, error) {
	c, err := req.Cookie(SessionCookie)
	if err == nil {
		return c.Value, err
	}
	return "", err
}

func (s *SessionHTTP) Save(w http.ResponseWriter) (string, time.Time) {
	sessionToken, expires := session.GenerateToken()
	http.SetCookie(w, &http.Cookie{
		Name:    SessionCookie,
		Value:   sessionToken,
		Expires: expires,
	})
	return sessionToken, expires
}

func (s *SessionHTTP) Clear(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   SessionCookie,
		MaxAge: -1,
	})
}
