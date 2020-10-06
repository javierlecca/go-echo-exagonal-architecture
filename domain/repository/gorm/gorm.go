package gorm

import (
	"test-github/domain/repository"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
)

type Gorm struct {
	c *repository.Config
}

func (g *Gorm) NewDb() []*gorm.DB {
	g.c = &repository.Config{}

	c, _ := g.c.NewConfig()

	arrayConnections := make([]*gorm.DB, 0)

	var db *gorm.DB

	var err error

	if len(c.DB.Mysql) > 0 {
		for _, v := range c.DB.Mysql {
			psqlConnStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				v.UserName,
				v.Password,
				v.Host,
				v.Port,
				v.Database)

			db, err = gorm.Open("mysql", psqlConnStr)

			if err != nil {
				fmt.Printf("err mysql = %v \n", err)
			}

			arrayConnections = append(arrayConnections, db)

			db = &gorm.DB{}
		}
	}
	return arrayConnections
}
