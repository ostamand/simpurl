package mysql

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/stretchr/testify/assert"
)

var l = &store.LinkModel{
	ID:          1,
	Symbol:      "robots",
	URL:         "https://www.google.com/robots.txt",
	Description: "Google manage crawler traffic",
	Note:        "My notest",
}

// reference: https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	return db, mock
}

func TestFindBySymbol(t *testing.T) {
	db, mock := NewMock()

	store := &linkSQL{db}
	defer db.Close()

	query := "SELECT id, symbol, url, description, note, created_at FROM links WHERE symbol = \\?"

	rows := sqlmock.NewRows([]string{"id", "symbol", "url", "description", "note", "created_at"}).AddRow(l.ID, l.Symbol, l.URL, l.Description, l.Note, time.Now())
	mock.ExpectQuery(query).WithArgs(l.Symbol).WillReturnRows(rows)

	linkFromSymbol, _ := store.FindBySymbol(l.Symbol)
	assert.Equal(t, l.URL, linkFromSymbol.URL)
}
