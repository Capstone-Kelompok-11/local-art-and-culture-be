package routes

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) *echo.Echo  {
	godotenv.Load("")
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))
	RoleRoute(e, db)
	AdminRoute(e, db)
	ArticleRoute(e, db)
	UserRoute(e, db)
	CreatorRoute(e, db)
	CategoryRoute(e, db)
	ProductRoute(e, db)
	ShippingRoute(e, db)
	EventRoute(e, db)
	CommentRoute(e, db)
	PaymentRoute(e, db)
	LikeRoute(e, db)
	CommentRoute(e, db)
	TicketRoute(e, db)
	GuestRoute(e, db)
	RatingRoute(e, db)
	WishlistRoute(e, db)
	TransactionRoute(e, db)
	TransactionDetailRoute(e, db)
	FilesRoute(e, db)
	ChatbotRoute(e, db)
	AuthGoogleRoute(e, db)
	return e
}