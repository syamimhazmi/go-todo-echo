package model

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
	ID   int    `json:"id" query:"id"`
	Task string `json:"task" query:"task"`
	Done bool   `json:"done" query:"done"`
}

func LoadDB() {
	var err error

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()
	//
	// err = db.Ping()
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }

	database.MigrateTables(db)
}

func GetTodos() ([]Todo, error) {
	rows, err := db.Query("select id, task, done from todos order by id desc")

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

func GetTodoById(id int) (*Todo, error) {
	var todo Todo

	err := db.QueryRow("select id, task, done from todos where id = $1", id).Scan(&todo.ID, &todo.Task, &todo.Done)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("todo with id %d not found", id)
		}

		return nil, err
	}

	return &todo, nil
}

func AddTodo(todo *Todo) error {
	err := db.QueryRow("insert into todos(task, done) values ($1, $2) returning id", todo.Task, todo.Done).Scan(&todo.ID)

	return err
}

func UpdateTodo(todo *Todo) error {
	err := db.QueryRow(
		`update todos set task = $1, done = $2 where id = $3 returning id, task, done`,
		todo.Task, todo.Done, todo.ID).Scan(&todo.ID, &todo.Task, &todo.Done)

	return err
}

func DeleteTodo(id int) error {
	_, err := db.Exec("delete from todos where id = $1", id)

	return err
}
