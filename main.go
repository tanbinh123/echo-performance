package main

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/echo-performance/common"
	"github.com/tomoyane/echo-performance/controller"
)

func main() {
	common.InitDB()
	e := echo.New()

	e.GET("/", controller.HelloController)
	e.GET("/items", controller.ItemController)

	e.Logger.Fatal(e.Start(":8080"))
}


