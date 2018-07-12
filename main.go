package main

import (
	"github.com/labstack/echo"
	"github.com/tomo0111/echo-performance/controller"
	"github.com/tomo0111/echo-performance/common"
)

func main() {
	common.InitDB()
	e := echo.New()

	e.GET("/", controller.HelloController)
	e.GET("/items", controller.ItemController)

	e.Logger.Fatal(e.Start(":8080"))
}


