package handlers

import (
	"log"
	"main/services"
	"main/templates"
	"os"

	"github.com/labstack/echo/v4"
)

func GetAbout(ctx echo.Context) error {
	about, err := os.ReadFile("./public/about.md")

	if err != nil {
		log.Println("Failed to read about file")
	}

	rendered := services.MdToHTML(about)

	return services.RenderPage(ctx, templates.Article(string(rendered)))
}
