package store

import "time"

//TODO refactor this StorageService to include seperate subservices

type LinkStorage interface {
	GetAll(u *UserModel) (*[]LinkModel, error)
	FindBySymbol(symbol string) (*LinkModel, error)
	Save(l *LinkModel) error
}

type UserStorage interface {
	Save(u *UserModel) error
	Delete(id int) error
	GetBySession(token string) (*UserModel, error)
	GetByUsername(username string) (*UserModel, error)
	DeleteFromUsername(username string) error
}

type SessionStorage interface {
	Save(session *SessionModel) error
	DeleteFromToken(token string) error
}

type Storage interface {
	Close()
}

type StorageService struct {
	Link    LinkStorage
	User    UserStorage
	Session SessionStorage
	Storage
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
	Note        string
	CreatedAt   time.Time
}

type UserModel struct {
	ID             int
	Username       string
	HashedPassword string
	Password       string
	Admin          bool
	CreatedAt      time.Time
}

func (u UserModel) Authenticated() bool {
	return u.ID != 0
}