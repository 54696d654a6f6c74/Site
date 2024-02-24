package handlers

import (
	articles "main/domain"
	"main/services"
	"main/templates"

	"github.com/a-h/templ"

	"github.com/labstack/echo/v4"
)

func renderArticles() []templ.Component {
	renderedArticles := []templ.Component{}

	for _, articleContent := range articles.GetArticles() {
		rendered := string(services.MdToHTML(articleContent))

		renderedArticles = append(renderedArticles, templates.Article(rendered))
	}

	return renderedArticles
}

func Index(ctx echo.Context) error {
	return services.RenderPage(ctx, templates.Index(renderArticles()[:2]))
}
