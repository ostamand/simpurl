package mysql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ostamand/url/web/store"
	"github.com/stretchr/testify/assert"
)

var l = &store.LinkModel{
	ID:          1,
	Symbol:      "robots",
	URL:         "https://www.google.com/robots.txt",
	Description: "Google manage crawler traffic",
}

// reference: https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	return db, mock
}

func TestFindBySymbol(t *testing.T) {
	db, mock := NewMock()

	store := &storageSQL{db}
	defer store.Close()

	query := "SELECT id, symbol, url, description FROM links WHERE symbol = \\?"

	rows := sqlmock.NewRows([]string{"id", "symbol", "url", "description"}).AddRow(l.ID, l.Symbol, l.URL, l.Description)
	mock.ExpectQuery(query).WithArgs(l.Symbol).WillReturnRows(rows)

	linkFromSymbol, _ := store.FindBySymbol(l.Symbol)
	assert.Equal(t, l.URL, linkFromSymbol.URL)
}
