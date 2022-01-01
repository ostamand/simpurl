package mysql

import (
	"database/sql"
	"log"
	"time"

	"github.com/ostamand/url/web/store"
	"golang.org/x/crypto/bcrypt"
)

type userSQL struct {
	db *sql.DB
}

func (storage userSQL) Save(username string, password string) error {
	var hashedPassword []byte
	var err error
	// TODO move this
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 8); err != nil {
		return err
	}
	// by default admin will be false
	stmt, _ := storage.db.Prepare("INSERT INTO users(username, password, created_at) VALUES(?, ?, ?)")
	defer stmt.Close()
	_, err = stmt.Exec(username, hashedPassword, time.Now())
	return err
}

func (storage userSQL) GetByUsername(username string) (*store.UserModel, error) {
	query := "SELECT id, username, password, admin, created_at FROM users WHERE username = ?"
	u := &store.UserModel{}
	err := storage.db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password, &u.Admin, &u.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return u, err
}

func (storage userSQL) GetBySession(token string) (*store.UserModel, error) {
	query := `SELECT users.id, username, password, users.admin, users.created_at from sessions 
	JOIN users ON sessions.user_id = users.id 
	WHERE token = ? AND expiry_at > ? 
	ORDER BY sessions.expiry_at DESC LIMIT 1`
	u := &store.UserModel{}
	err := storage.db.QueryRow(query, token, time.Now()).Scan(&u.ID, &u.Username, &u.Password, &u.Admin, &u.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return u, err
}
