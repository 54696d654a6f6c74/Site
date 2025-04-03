package main

import (
	"main/handlers"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", handlers.Index)
	e.GET("/about", handlers.GetAbout)
	e.GET("/articles", handlers.GetArticles)
	e.GET("/articles/:name", handlers.GetArticle)

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.GET("/data/articles", handlers.IndexArticles)

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
