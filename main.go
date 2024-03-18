package main

import (
	"log"
	"main/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", handlers.Index)
	e.GET("/about", handlers.GetAbout)
	e.GET("/articles", handlers.GetArticles)
	e.GET("/articles/:name", handlers.GetArticle)

	e.GET("/data/articles", handlers.IndexArticles)

	log.Println("Listening on port 3000")
	e.Logger.Fatal(e.Start(":3000"))
}
