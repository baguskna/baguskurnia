package controller

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

type HomeData struct {
	Title   string
	Content template.HTML
	Locale  string
}

func (h *HomeController) HomeHandler(c echo.Context) error {
	query := c.FormValue("hl")

	data := HomeData{
		Title: "👋 Hi, I'm Bagus Kurnia",
		Content: template.HTML(`
			<p class="text">Software engineer with 4 years of experiences</p>
			<p class="text">
				<span>
					Text stack
					<pre>.ts</pre>
					<pre>.go</pre>
				</span>
			</p>
		`),
		Locale: "en",
	}

	if query == "en" {
		return c.Render(200, "index", data)
	}

	if query == "ja" {
		data = HomeData{
			Title:   "👋 Hi, バグス・クルニアです",
			Content: template.HTML(`<p class="text"> 4年の経験があるソフトウェアエンジニア</p>`),
			Locale:  "ja",
		}
	}

	// handle not found query 'hl'
	if query != "" && query != "ja" {
		return c.Render(200, "404", "")
	}
	return c.Render(200, "index", data)
}
