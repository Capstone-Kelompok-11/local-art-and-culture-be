package migrations

import (
	"lokasani/entity/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.SuperAdmin{})
}
