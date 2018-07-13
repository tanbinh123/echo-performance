package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/tomo0111/echo-performance/common"
	"github.com/tomo0111/echo-performance/entity"
)

func ItemController (c echo.Context) error {
	var items []entity.Item
	return c.JSON(http.StatusOK, common.Db.Find(&items))
}