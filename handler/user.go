package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/s-frick/crypto-front/view"
	"github.com/s-frick/crypto-front/view/user"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return view.Render(c, user.Show())
}
