package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

var todos []Todo

var idCounter int

func main() {
	e := echo.New()

	renderer := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = renderer

	// enable logger so that i can see error log on terminal!
	// should've write on their documentation
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handleIndex)
	e.POST("/todos", handleCreateTodo)
	e.PUT("/todos/:id", handleToggleTodo)
	e.DELETE("/todos/:id", handleDeleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
	// e.GET("/todos/:id", handleEditTodo)

	// e.Logger.Fatal(e.Start(":5001"))
}

func handleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", todos)
}

func handleCreateTodo(c echo.Context) error {
	task := c.FormValue("task")

	idCounter++

	todo := Todo{
		ID:     idCounter,
		Task:   task,
		IsDone: false,
	}

	todos = append(todos, todo)

	return c.Render(http.StatusCreated, "todo-item", todos)
}

func handleToggleTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].IsDone = !todos[i].IsDone

			return c.Render(http.StatusOK, "todo-item", todos)
		}
	}

	return c.NoContent(http.StatusNotFound)
}
