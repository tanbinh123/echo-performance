package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloController (c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"key": "hello world"})
}