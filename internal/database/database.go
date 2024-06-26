package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Db *sql.DB
	Tx *sql.Tx
}

func (db *DB) InitDb() {
	connectionStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	db.Db = conn
	fmt.Println("Connected to database")

	conn.Close()
}
