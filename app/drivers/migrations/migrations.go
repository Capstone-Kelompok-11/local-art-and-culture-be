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
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Shipping{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Transaction{})
	db.AutoMigrate(&models.TransactionDetail{})
	db.AutoMigrate(&models.Like{})
	db.AutoMigrate(&models.Ticket{})
	db.AutoMigrate(&models.Wishlist{})
}
