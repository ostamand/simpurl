package session

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const ExpirationDelay = time.Minute * 60 * 3

type SessionClient interface {
	Get(req *http.Request) (string, error)
	Save(w http.ResponseWriter) (string, time.Time)
	Clear(w http.ResponseWriter)
}

func GenerateToken() (token string, expires time.Time) {
	token = uuid.NewV4().String()
	expires = time.Now().Add(ExpirationDelay)
	return
}