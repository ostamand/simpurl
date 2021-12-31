package mysql

import (
	"database/sql"
	"log"

	"github.com/ostamand/url/web/store"
)

type sessionSQL struct {
	db *sql.DB
}

func (storage sessionSQL) Save(sesssion *store.SessionModel) error {
	stmt, _ := storage.db.Prepare("INSERT INTO sessions(user_id, token, created_at, expiry_at) values(?, ?, ?, ?)")
	defer stmt.Close()
	_, err := stmt.Exec(sesssion.UserID, sesssion.Token, sesssion.CreatedAt, sesssion.ExpiryAt)
	if err != nil {
		log.Println(err)
	}
	return err
}
