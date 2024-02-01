package main

import (
	"fmt"
	"main/handlers"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()

	e.Static("/public", "public");

	e.GET("/", handlers.Index)

	fmt.Println("Listening on port 3000")
	e.Logger.Fatal(e.Start(":3000"))
}
