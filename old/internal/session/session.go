package session

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const ExpirationDelay = time.Minute * 60 * 3

func GenerateToken() (token string, expires time.Time) {
	token = uuid.NewV4().String()
	expires = time.Now().Add(ExpirationDelay)
	return
}