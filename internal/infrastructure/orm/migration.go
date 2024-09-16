package orm

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&Account{},
		&User{},
		&Allocation{},
		&Transaction{},
		&RawTransaction{},
	)
}
