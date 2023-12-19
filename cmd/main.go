package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/s-frick/crypto-front/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.Secure())

	userHandler := handler.UserHandler{}
	orderbookHandler := handler.OrderbookHandler{}
	e.GET("/", userHandler.HandleUserShow)
	e.GET("/orderbook", orderbookHandler.OrderbookShow)
	e.File("/styles/main.css", "assets/css/out.css")

	fmt.Print(e.Start(":3000"))

}
