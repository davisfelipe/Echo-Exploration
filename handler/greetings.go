package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	fmt.Println("Hola mundo")
	qParams := c.Param("q")
	if qParams == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "hola"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "hi"})
}
