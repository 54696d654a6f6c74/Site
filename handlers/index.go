package handlers

import (
	"fmt"
	"main/services"
	"main/templates"
	"os"
	"sort"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func getArticles() []templ.Component {
	articleFiles, err := os.ReadDir("./articles")

	if err != nil {
		fmt.Println(err)
	}

	sort.Slice(articleFiles, func(a, b int) bool {
		aInfo, err := articleFiles[a].Info()

		if err != nil {
			fmt.Println(err)
		}

		bInfo, err := articleFiles[b].Info()

		aTime := services.GetFileCreationTime(aInfo)
		bTime := services.GetFileCreationTime(bInfo)

		return aTime.Unix() > bTime.Unix()
	})

	articles := []templ.Component{}

	for _, file := range articleFiles {
		conent, err := os.ReadFile("./articles/" + file.Name())

		if err != nil {
			fmt.Println(err)
		}

		rendered := string(services.MdToHTML(conent))

		articles = append(articles, templates.Article(rendered))
	}

	return articles
}

func Index(ctx echo.Context) error {
	return renderPage(ctx, templates.Index(getArticles()))
}
