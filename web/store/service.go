package store

import "time"

//TODO refactor this StorageService to include seperate subservices
type StorageService interface {
	SaveLink(l *LinkModel) error
	SaveUser(username string, password string) error
	SaveSession(session *SessionModel) error
	GetUserBySession(token string) (*UserModel, error)
	GetByUsername(username string) (*UserModel, error)
	GetAllLinks(u *UserModel) (*[]LinkModel, error)
	FindBySymbol(symbol string) (*LinkModel, error)
	Close()
}

type SessionModel struct {
	ID        int
	UserID    int
	Token     string
	ExpiryAt  time.Time
	CreatedAt time.Time
}

type LinkModel struct {
	ID          int
	UserID      int
	Symbol      string
	URL         string
	Description string
	CreatedAt   time.Time
}

type UserModel struct {
	ID        int
	Username  string
	Password  string
	Admin     bool
	CreatedAt time.Time
}

func (u UserModel) Authenticated() bool {
	return u.ID != 0
}
