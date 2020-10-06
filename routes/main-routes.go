package routes

import (
	"test-github/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	c := controller.Controller{}

	// create groups
	apiGroup := e.Group("/api")
	// set group routes
	c.ApiGroup(apiGroup)

	return e
}
