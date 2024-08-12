package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"todo-echo/internals/model"
	"todo-echo/internals/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.LoadDB()

	e := echo.New()

	renderer := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = renderer

	// enable logger so that i can see error log on terminal!
	// should've write on their documentation
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
