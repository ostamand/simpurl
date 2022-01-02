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

func (storage userSQL) Save(u *store.UserModel) error {
	var hashedPassword []byte
	var err error
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(u.Password), 8); err != nil {
		return err
	}
	stmt, _ := storage.db.Prepare("INSERT INTO users(username, hashed_password, admin, created_at) VALUES(?, ?, ?, ?)")
	defer stmt.Close()
	_, err = stmt.Exec(u.Username, hashedPassword, u.Admin, time.Now())
	return err
}

func (storage userSQL) Delete(id int) error {
	stmt, _ := storage.db.Prepare("DELETE FROM users WHERE id = ?")
	defer stmt.Close()
	_, err := stmt.Exec(id)
	return err
}

func (storage userSQL) GetByUsername(username string) (*store.UserModel, error) {
	query := "SELECT id, username, hashed_password, admin, created_at FROM users WHERE username = ?"
	u := &store.UserModel{}
	err := storage.db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.HashedPassword, &u.Admin, &u.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return u, err
}

// TODO refactor this. won't work for 2 users
func (storage userSQL) GetBySession(token string) (*store.UserModel, error) {
	query := `SELECT users.id, username, hashed_password, users.admin, users.created_at from sessions 
	JOIN users ON sessions.user_id = users.id 
	WHERE token = ? AND expiry_at >= ? 
	ORDER BY sessions.expiry_at DESC LIMIT 1`
	u := &store.UserModel{}
	err := storage.db.QueryRow(query, token, time.Now()).Scan(&u.ID, &u.Username, &u.HashedPassword, &u.Admin, &u.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return u, err
}
