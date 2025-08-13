package dao

import "gorm.io/gorm"

func Inittable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
