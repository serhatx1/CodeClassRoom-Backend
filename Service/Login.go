package Service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}
