package view

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, c templ.Component) error {
	return c.Render(ctx.Request().Context(), ctx.Response())
}
