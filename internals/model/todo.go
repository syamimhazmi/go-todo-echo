package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo-echo/internals/database"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func init() {
	var err error

	fmt.Printf("database url: %s\n, app port: %s\n", os.Getenv("DATABASE_URL"), os.Getenv("APP_PORT"))

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	database.MigrateTables(db)
}

func GetTodos() ([]Todo, error) {
	rows, err := db.Query("select id, task, done from todos")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.ID, &todo.Task, &todo.Done)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func AddTodo(todo *Todo) error {
	err := db.QueryRow("insert into todos(title, done) values ($1, $1) returning id", todo.Task, todo.Done).Scan(&todo.ID)

	return err
}

func UpdateTodo(todo *Todo) error {
	_, err := db.Exec("update todos set task = $1, done = $2 where id = $3", todo.Task, todo.Done, todo.ID)

	return err
}

func DeleteTodo(id int) error {
	_, err := db.Exec("delete from todos where id = $1", id)

	return err
}
