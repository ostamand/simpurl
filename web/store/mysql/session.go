package mysql

import (
	"database/sql"
	"log"

	"github.com/ostamand/url/web/store"
)

type sessionSQL struct {
	db *sql.DB
}

func (storage *sessionSQL) Save(session *store.SessionModel) error {
	stmt, _ := storage.db.Prepare("INSERT INTO sessions(user_id, token, created_at, expiry_at) values(?, ?, ?, ?)")
	defer stmt.Close()
	_, err := stmt.Exec(session.UserID, session.Token, session.CreatedAt, session.ExpiryAt)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (storage *sessionSQL) DeleteFromToken(token string) error {
	stmt, _ := storage.db.Prepare("DELETE FROM sessions WHERE token = ?")
	defer stmt.Close()
	_, err := stmt.Exec(token)
	return err
}
