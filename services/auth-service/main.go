package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		fmt.Println("authorization header", c.Request().Header.Get("authorization"))
		return c.JSON(200, map[string]string{
			"service": "auth-service",
		})
	})
	server.Start(":8080")
}
