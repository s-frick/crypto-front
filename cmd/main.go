package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/s-frick/crypto-front/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := handler.UserHandler{}
	e.GET("/", userHandler.HandleUserShow)

	e.Start(":3000")

}
