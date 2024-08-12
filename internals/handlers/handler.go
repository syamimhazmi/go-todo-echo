package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	models "todo-echo/internals/model"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	todos, err := models.GetTodos()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.Render(http.StatusOK, "index", todos)
}

func CreateTodo(c echo.Context) error {
	todo := new(models.Todo)

	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	todo.Task = c.FormValue("task")

	err := models.AddTodo(todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	todo := new(models.Todo)

	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	todo.ID = id

	err = models.UpdateTodo(todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = models.DeleteTodo(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
