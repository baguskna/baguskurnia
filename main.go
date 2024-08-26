package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &Templates{
		templates: template.Must(template.ParseGlob(dir + "/views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", "")
	})

	e.GET("/test", func(c echo.Context) error {
		return c.Render(200, "test", "")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
