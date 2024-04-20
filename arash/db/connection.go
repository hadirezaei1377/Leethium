package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {

	const (
		host     = "localhost"
		port     = 5432
		user     = "username"
		password = "password"
		dbname   = "database_name"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the PostgreSQL database")

	return db, nil
}
