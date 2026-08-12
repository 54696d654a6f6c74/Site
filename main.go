package main

import (
	"fmt"
	"os"

	"main/handlers"

	"github.com/joho/godotenv"
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

	godotenv.Load()
	envPort, ok := os.LookupEnv("PORT")

	var port string

	if ok {
		port = envPort
	} else {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
