package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloController (c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

