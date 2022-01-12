package test

import (
	"errors"
	"net/http"
	"time"

	"github.com/ostamand/simpurl/internal/session"
)

type SessionMock struct {
	Token string
}

func (s *SessionMock) Save(http.ResponseWriter) (string, time.Time) {
	sessionToken, expires := session.GenerateToken()
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