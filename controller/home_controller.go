package controller

import (
	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

type HomeData struct {
	Title   string
	Content string
}

func (h *HomeController) HomeHandler(c echo.Context) error {
	query := c.FormValue("hl")

	data := HomeData{
		Title:   "👋 Hi, I'm Bagus Kurnia",
		Content: "Software engineer with 4 years of experiene",
	}

	if query == "jp" {
		data = HomeData{
			Title:   "👋 Hi, バグスです",
			Content: "Software engineer with 4 years of experiene",
		}
	}

	return c.Render(200, "index", data)
}
