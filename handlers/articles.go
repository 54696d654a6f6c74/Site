package handlers

import (
	articles "main/domain/article"
	"main/services"
	"main/templates"

	"github.com/labstack/echo/v4"
)

func GetArticles(ctx echo.Context) error {
	return services.RenderPage(ctx, templates.ArticleListPage(articles.GetArticles()))
}
