package services

import (
	"main/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderComponent(ctx echo.Context, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(200)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func RenderPage(ctx echo.Context, page templ.Component) error {
	return RenderComponent(ctx, templates.Wrapper(page))
}
