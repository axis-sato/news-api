package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	devDataSourceName = "news_crawler:news_crawler_password@tcp(localhost:3309)/news_crawler?charset=utf8mb4&parseTime=True&loc=Local"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", devDataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}