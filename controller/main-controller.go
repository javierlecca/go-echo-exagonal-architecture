package controller

import (
	"test-github/domain/service/society"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	society *society.SocietyService
}

func (c *Controller) NewsocietyController() {
	c.society = &society.SocietyService{}
	c.society.NewsocietyOrm()
}

func (c *Controller) ApiGroup(g *echo.Group) {
	c.NewsocietyController()
	g.GET("/societies", c.society.ListAll)
}
