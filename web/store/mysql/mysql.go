package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ostamand/url/web/config"
	"github.com/ostamand/url/web/store"
)

type storageSQL struct {
	db *sql.DB
}

func (s storageSQL) Close() {
	s.db.Close()
}

func InitializeSQL(params *config.ParamsDB) *store.StorageService {
	dataSourceName := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?parseTime=true",
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

	s := store.StorageService{
		Link:    linkSQL{db},
		User:    userSQL{db},
		Session: &sessionSQL{db},
		Storage: storageSQL{db},
	}

	return &s
}
