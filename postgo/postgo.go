package postgo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username"
	password = "password"
	dbname   = "dbname"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	readData(db)

	insertData(db)
}

func readData(db *sql.DB) {
	query := `SELECT column_name FROM table_name WHERE condition;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			log.Fatal(err)
		}
		fmt.Println(columnName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB) {
	query := `INSERT INTO table_name (column1, column2) VALUES ($1, $2)`
	_, err := db.Exec(query, "value1", "value2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully!")
}
