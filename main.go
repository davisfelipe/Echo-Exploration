package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/davisfelipe/Echo-Exploration/handler"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func greeting(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hola %s\n", name)
}

func main() {
	e := echo.New()

	// Extra configuration 3-parties
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// Config Middleware Functions
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Config routes information
	e.GET("/", func(c echo.Context) error {
		var wg sync.WaitGroup

		wg.Add(2)
		go greeting("pedro", &wg)
		go greeting("pablo", &wg)

		wg.Wait()

		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello", handler.Hello)
	e.Logger.Fatal(e.Start(":5000"))
}
