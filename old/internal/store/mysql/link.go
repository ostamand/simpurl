package mysql

import (
	"database/sql"
	"log"
	"time"

	"github.com/ostamand/simpurl/internal/store"
)

type linkSQL struct {
	db *sql.DB
}

func (storage *linkSQL) GetAll(userID int) (*[]store.LinkModel, error) {
	var links []store.LinkModel

	query := "SELECT id, symbol, url, description, note, created_at FROM links WHERE user_id = ?"

	stmt, _ := storage.db.Prepare(query)
	defer stmt.Close()

	rows, _ := stmt.Query(userID)
	defer rows.Close()

	for rows.Next() {
		l := store.LinkModel{UserID: userID}
		rows.Scan(&l.ID, &l.Symbol, &l.URL, &l.Description, &l.Note, &l.CreatedAt)
		links = append(links, l)
	}
	return &links, rows.Err()
}

func (storage *linkSQL) FindBySymbol(userID int, symbol string) (*store.LinkModel, error) {
	var l store.LinkModel
	query := "SELECT id, symbol, url, description, note, created_at FROM links WHERE symbol = ? AND user_id = ?"
	err := storage.db.QueryRow(query, symbol, userID).Scan(&l.ID, &l.Symbol, &l.URL, &l.Description, &l.Note, &l.CreatedAt)
	return &l, err
}

func (storage *linkSQL) Save(l *store.LinkModel) error {
	stmt, _ := storage.db.Prepare("INSERT INTO links(user_id, symbol, url, description, note, created_at) values(?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	_, err := stmt.Exec(l.UserID, l.Symbol, l.URL, l.Description, l.Note, time.Now())

	if err != nil {
		log.Println(err)
	}
	return err
}

func (storage *linkSQL) DeleteByURL(url string) error {
	stmt, _ := storage.db.Prepare("DELETE FROM links WHERE url = ?")
	defer stmt.Close()
	_, err := stmt.Exec(url)
	return err
}
