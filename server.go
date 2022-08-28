package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var upgarder = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartServer(c echo.Context) error {
	ws, err := upgarder.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	id := c.Param("id")
	fmt.Printf("%s", id)

	for {
		err := ws.WriteJSON(map[string]interface{}{"title": "Ant Design Title"})
		// err := ws.WriteMessage(websocket.TextMessage, []byte("Hello World!"))
		if err != nil {
			c.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", StartServer)
	e.Logger.Fatal(e.Start(":1323"))
}
