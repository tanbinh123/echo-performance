package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func ItemController (c echo.Context) error {
	return c.String(http.StatusOK, "Item")
}

