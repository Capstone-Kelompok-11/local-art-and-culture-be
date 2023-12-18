package migrations

import (
	"lokasani/entity/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Transaction{})
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
	db.AutoMigrate(&models.TransactionDetail{})
	db.AutoMigrate(&models.Like{})
	db.AutoMigrate(&models.Ticket{})
	db.AutoMigrate(&models.Guest{})
	db.AutoMigrate(&models.Rating{})
	db.AutoMigrate(&models.Wishlist{})
	db.AutoMigrate(&models.Files{})
	db.AutoMigrate(&models.SaveChatbot{})
	db.Migrator().AddColumn(&models.Creator{}, "Address")
	db.Migrator().AddColumn(&models.Product{}, "Stock")
	db.Migrator().AddColumn(&models.Product{}, "TotalProduct")
}
