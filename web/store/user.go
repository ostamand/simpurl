package store

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const ExpirationDelay = time.Minute * 60 * 3

func CreateSession(storage *StorageService, user *UserModel) (*SessionModel, error) {
	sessionToken := uuid.NewV4().String()
	session := SessionModel{
		UserID:    user.ID,
		Token:     sessionToken,
		CreatedAt: time.Now(),
		ExpiryAt:  time.Now().Add(ExpirationDelay),
	}
	err := (*storage).SaveSession(&session)
	return &session, err
}

func VerifyPassword(storage *StorageService, userName string, password string) (*UserModel, error) {
	user, err := (*storage).GetByUsername(userName)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
