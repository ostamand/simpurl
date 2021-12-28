package store

import "time"

type StorageService interface {
	SaveLink(l *LinkModel) error
	SaveUser(userName string, password string) error
	SaveSession(session *SessionModel) error
	GetByUsername(userName string) (*UserModel, error)
	FindBySymbol(symbol string) (*LinkModel, error)
	Close()
}

type SessionModel struct {
	ID        int
	UserID    int
	Token     string
	CreatedAt time.Time
	ExpiryAt  time.Time
}

type LinkModel struct {
	ID          int
	Symbol      string
	URL         string
	Description string
}

type UserModel struct {
	ID       int
	Username string
	Password string
}
