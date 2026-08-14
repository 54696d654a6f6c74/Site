package handlers

import (
	"log"
	articles "main/domain/article"
	"main/services"
	"main/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type params struct {
	Name string `param:"name"`
}

func getArticle(name string) templ.Component {
	article := articles.GetArticle(name)

	if article == nil {
		return nil
	}

	return templates.Article(article)
}

func GetArticle(ctx echo.Context) error {
	params := params{}

	err := (&echo.DefaultBinder{}).BindPathParams(ctx, &params)

	if err != nil {
		log.Println("Failed to parse path params")
	}

	pageContent := getArticle(params.Name)
	if pageContent == nil {
		return echo.NewHTTPError(404, "Article does not exist")
	}

	return services.RenderPage(ctx, pageContent)
}
