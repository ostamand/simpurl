package store

type StorageService interface {
	SaveLink(l *LinkModel) error
	SaveUser(userName string, password string) error
	FindBySymbol(symbol string) (*LinkModel, error)
	Close()
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
