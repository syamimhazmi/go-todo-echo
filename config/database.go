package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() {
	var err error

	db, err := sql.Open("postgres", "postgresql://postgres@localhost:5432/todos?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	migrateTables(db)
}

func migrateTables(db *sql.DB) {
	createTodosTable(db)
}

func createTodosTable(db *sql.DB) {
	_, err := db.Exec(`
        create table if not exists todos (
            id serial primary key,
            task text not null,
            done boolean not null default false
        )
    `)

	if err != nil {
		log.Fatal(err)
	}
}
