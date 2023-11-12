package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func Hello(c echo.Context) error {
	fmt.Println("Hola mundo")
	qParams := c.Param("q")
	if qParams == "" {
		return c.JSON(http.StatusBadRequest, HelloResponse{
			Message: "Hola Mundo",
			Success: true,
		})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "hi"})
}
