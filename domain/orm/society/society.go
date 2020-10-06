package society

import (
	"errors"

	g_ "test-github/domain/repository/gorm"

	"github.com/jinzhu/gorm"

	"test-github/domain/entity/society"
)

// db, _ := gorm.NewDb()

type SocietyOrm struct {
	db *gorm.DB
}

func (s *SocietyOrm) NewsocietyOrm() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]

}

func (s *SocietyOrm) GetAll() ([]*society.Society, error) {

	societies := make([]*society.Society, 0)

	if s.db.Find(&societies).Error != nil {
		return nil, errors.New("Datacenter not found")
	}

	return societies, nil
}
