package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(dsn string) *Queries {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return New(conn)
}
