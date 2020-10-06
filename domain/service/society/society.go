package society

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"test-github/domain/orm/society"
	"test-github/util"
)

type SocietyService struct {
	orm  *society.SocietyOrm
	resp *util.Response
}

func (s *SocietyService) NewsocietyOrm() {
	s.resp = &util.Response{}
	s.orm = &society.SocietyOrm{}
	s.orm.NewsocietyOrm()
}

func (s *SocietyService) ListAll(c echo.Context) error {
	societies, _ := s.orm.GetAll()
	s.resp.Status = "OK"
	s.resp.Data = societies
	return c.JSON(http.StatusOK, s.resp)
}
