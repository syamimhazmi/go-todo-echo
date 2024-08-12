package database

import (
	"database/sql"
	"log"
)

func MigrateTables(db *sql.DB) {
	createTodoTables(db)
}

func createTodoTables(db *sql.DB) {
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
