package migrations

import (
	"lokasani/entity/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.SuperAdmin{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Creator{})
	db.AutoMigrate(&models.Category{})
}
