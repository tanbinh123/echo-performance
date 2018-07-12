package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomo0111/echo-performance/common"
	"github.com/tomo0111/echo-performance/entity"
)

func ItemController (c echo.Context) error {
	items := common.Db.Find([]entity.Item{})
	return c.JSON(http.StatusOK, items)
}