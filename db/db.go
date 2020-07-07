package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	c, err := readDBConf()

	if err != nil {
		return nil, err
	}

	return sql.Open(c.Dialect, c.Datasource)
}
