package dataprocessing

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM users WHERE age > 30 ORDER BY age DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id   int
		name string
		age  int
	)

	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, age)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
