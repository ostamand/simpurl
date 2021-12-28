package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ostamand/url/web/config"
	"github.com/ostamand/url/web/store"
)

type storageSQL struct {
	db *sql.DB
}

func InitializeSQL(params *config.ParamsDB) *storageSQL {
	dataSourceName := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		params.User, params.Pass, params.Addr, params.Port, params.Name,
	)

	// connect with unix sockets for cloud run
	// ref: https://cloud.google.com/sql/docs/mysql/connect-run#public-ip-default
	if params.Instance != "" && params.SocketDir != "" {
		dataSourceName = fmt.Sprintf(
			"%s:%s@unix(/%s/%s)/%s?parseTime=true",
			params.User, params.Pass, params.SocketDir, params.Instance, params.Name,
		)
	}

	log.Printf("connecting to: %s\n", dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return &storageSQL{db}
}

func (s storageSQL) Close() {
	s.db.Close()
}

func (s storageSQL) FindBySymbol(symbol string) (*store.LinkModel, error) {
	var l store.LinkModel
	query := "SELECT id, symbol, url, description FROM links WHERE symbol = ?"
	err := s.db.QueryRow(query, symbol).Scan(&l.ID, &l.Symbol, &l.URL, &l.Description)
	return &l, err
}

func (s storageSQL) SaveLink(l *store.LinkModel) error {
	stmt, _ := s.db.Prepare("INSERT INTO links(user_id, symbol, url, description) values(?, ?, ?, ?)")
	defer stmt.Close()
	_, err := stmt.Exec(l.UserID, l.Symbol, l.URL, l.Description)
	return err
}

func (s storageSQL) SaveUser(userName string, password string) error {
	var hashedPassword []byte
	var err error
	// TODO move this
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 8); err != nil {
		return err
	}
	log.Printf("usernam: %s password %s", userName, hashedPassword)
	stmt, _ := s.db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	defer stmt.Close()
	_, err = stmt.Exec(userName, hashedPassword)
	log.Println(err)
	return err
}

func (s storageSQL) SaveSession(sesssion *store.SessionModel) error {
	stmt, _ := s.db.Prepare("INSERT INTO sessions(user_id, token, created_at, expiry_at) values(?, ?, ?, ?)")
	defer stmt.Close()
	_, err := stmt.Exec(sesssion.UserID, sesssion.Token, sesssion.CreatedAt, sesssion.ExpiryAt)
	return err
}

func (s storageSQL) GetByUsername(userName string) (*store.UserModel, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	user := &store.UserModel{}
	err := s.db.QueryRow(query, userName).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func (s storageSQL) GetUserBySession(token string) (*store.UserModel, error) {
	query := `SELECT users.id, username, password from sessions 
	JOIN users ON sessions.user_id = users.id 
	WHERE token = ? AND expiry_at > ? 
	ORDER BY sessions.expiry_at DESC LIMIT 1`
	user := &store.UserModel{}
	err := s.db.QueryRow(query, token, time.Now()).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
