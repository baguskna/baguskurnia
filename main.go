package main

import (
	"baguskurnia/controller"
	"html/template"
	"io"
	"log"
	"net/http"
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

	homeController := controller.NewHomeController()
	// Custom 404 handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == http.StatusNotFound {
				c.Render(200, "404", "")
				return
			}
		}
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.GET("/", homeController.HomeHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
