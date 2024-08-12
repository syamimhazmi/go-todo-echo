package routes

import (
	"todo-echo/internals/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.GetTodos)

	e.POST("/todos", handlers.CreateTodo)

	e.PUT("/todos/:id", handlers.UpdateTodo)

	e.DELETE("/todos/:id", handlers.DeleteTodo)
}
