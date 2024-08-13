package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-echo/internals/model"
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

	todos, err := models.GetTodos()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.Render(http.StatusCreated, "todo-item", todos)
}

func EditTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	todo, err := model.GetTodoById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.HTML(http.StatusOK, fmt.Sprintf(`
		<input type="text" 
			   class="p-2 border rounded"
               name="task" 
               value="%s" 
               hx-put="/todos/%d" 
               hx-trigger="blur" 
               hx-target="#todo-%d" 
               hx-swap="outerHTML">
		`, todo.Task, id, id))
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
	todo.Task = c.FormValue("task")

	err = models.UpdateTodo(todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.HTML(http.StatusOK, fmt.Sprintf(`
        <li id="todo-%d" class="flex items-center justify-between p-2 border-b">
          <span class="cursor-pointer" 
                hx-get="/todos/%d/edit" 
                hx-trigger="click" 
                hx-target="this" 
                hx-swap="outerHTML">%s</span>
          <input type="hidden" name="done" value="%v">
          <button hx-delete="/todos/%d" 
                  hx-target="#todo-%d" 
                  class="p-2 bg-red-500 text-white rounded">Delete</button>
        </li>
    `, id, id, todo.Task, todo.Done, id, id))
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

	todos, err := model.GetTodos()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.Render(http.StatusOK, "todo-item", todos)
}
