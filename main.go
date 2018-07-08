package main

import (
	"github.com/labstack/echo"
	"github.com/tomo0111/echo-performance/controller"
)

func main() {
	e := echo.New()

	e.GET("/", controller.HelloController)
	e.GET("/items", controller.ItemController)

	e.Logger.Fatal(e.Start(":8080"))
}


