package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/tomoyane/echo-performance/entity"
	"github.com/tomoyane/echo-performance/common"
)

func ItemController (c echo.Context) error {
	var items []entity.Item
	common.Db.Find(&items)
	return c.JSON(http.StatusOK, &items)
}