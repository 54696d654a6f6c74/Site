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
	articleFolders, err := os.ReadDir("./articles")

	if err != nil {
		fmt.Println(err)
	}

	articleFolders = services.FilterForFolders(articleFolders)

	sort.Slice(articleFolders, func(a, b int) bool {
		aInfo, err := articleFolders[a].Info()

		if err != nil {
			fmt.Println(err)
		}

		bInfo, err := articleFolders[b].Info()

		aTime := services.GetFileChangeTime(aInfo)
		bTime := services.GetFileChangeTime(bInfo)

		return aTime.Unix() > bTime.Unix()
	})

	articles := []templ.Component{}

	for _, file := range articleFolders {
		conent, err := os.ReadFile("./articles/" + file.Name() + "/article.md")

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
