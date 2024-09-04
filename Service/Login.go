package Service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	fmt.Println("hello")
	return c.JSON(http.StatusOK, nil)
}
