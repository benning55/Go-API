package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well, Hello there")
	})
	e.Logger.Print("Listening to port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
