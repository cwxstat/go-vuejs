package main

import (
	"example.com/m/v2/handles"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Use this docker
	//e.Static("/", "../dist")

	e.Static("/", "./dist")

	e.GET("/ws", handles.Hello)
	e.GET("/status", handles.GetStatus)
	e.Logger.Fatal(e.Start(":8080"))
}
