package handlers

import (
	"log"
	articles "main/domain"
	"main/services"
	"main/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type params struct {
	Name string `param:"name"`
}

func getArticle(name string) templ.Component {
	rendered := string(services.MdToHTML(articles.GetArticle(name)))

	return templates.Article(rendered)
}

func GetArticle(ctx echo.Context) error {
	params := params{}

	err := (&echo.DefaultBinder{}).BindPathParams(ctx, &params)

	if err != nil {
		log.Println("Failed to parse path params")
	}

	return services.RenderPage(ctx, getArticle(params.Name))
}
