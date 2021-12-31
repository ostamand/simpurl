package mysql

import (
	"database/sql"
	"log"
	"time"

	"github.com/ostamand/url/web/store"
)

type linkSQL struct {
	db *sql.DB
}

func (storage linkSQL) GetAll(u *store.UserModel) (*[]store.LinkModel, error) {
	var links []store.LinkModel

	query := "SELECT id, symbol, url, description, note, created_at FROM links WHERE user_id = ?"

	stmt, _ := storage.db.Prepare(query)
	defer stmt.Close()

	rows, _ := stmt.Query(u.ID)
	defer rows.Close()

	for rows.Next() {
		l := store.LinkModel{UserID: u.ID}
		rows.Scan(&l.ID, &l.Symbol, &l.URL, &l.Description, &l.Note, &l.CreatedAt)
		links = append(links, l)
	}
	return &links, rows.Err()
}

func (storage linkSQL) FindBySymbol(symbol string) (*store.LinkModel, error) {
	var l store.LinkModel
	query := "SELECT id, symbol, url, description, note, created_at FROM links WHERE symbol = ?"
	err := storage.db.QueryRow(query, symbol).Scan(&l.ID, &l.Symbol, &l.URL, &l.Description, &l.Note, &l.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return &l, err
}

func (storage linkSQL) Save(l *store.LinkModel) error {
	stmt, _ := storage.db.Prepare("INSERT INTO links(user_id, symbol, url, description, note, created_at) values(?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	_, err := stmt.Exec(l.UserID, l.Symbol, l.URL, l.Description, l.Note, time.Now())
	return err
}
