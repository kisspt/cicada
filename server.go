package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	hub := newHub()
	go hub.run()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", func(c echo.Context) error {
		err := serveWs(hub, c)
		return err
	})
	e.Logger.Fatal(e.Start(":1323"))
}
