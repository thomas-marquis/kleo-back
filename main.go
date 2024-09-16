package main

import (
	"github.com/thomas-marquis/kleo-back/internal/infrastructure/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "admin:admin@tcp(localhost:3306)/kleo?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	orm.Migrate(db)
}
