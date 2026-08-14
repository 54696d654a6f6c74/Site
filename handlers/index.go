package handlers

import (
	articles "main/domain/article"
	"main/services"
	"main/templates"

	"github.com/a-h/templ"

	"github.com/labstack/echo/v4"
)

func renderArticles() []templ.Component {
	renderedArticles := []templ.Component{}

	for _, article := range articles.GetArticles() {
		renderedArticles = append(renderedArticles, templates.Article(article))
	}

	return renderedArticles
}

func IndexArticles(ctx echo.Context) error {
	indexArticles := renderArticles()

	if len(indexArticles) >= 2 {
		indexArticles = indexArticles[:2]
	}

	return services.RenderComponent(ctx, templates.IndexArticles(indexArticles))
}

func Index(ctx echo.Context) error {
	return services.RenderPage(ctx, templates.Index())
}
