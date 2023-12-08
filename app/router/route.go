package routes

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) *echo.Echo {
	godotenv.Load(".env")
	e := echo.New()
	e.Use(middleware.CORS())
	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eJwt.Use(middleware.CORS())
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
	ChatbotRoute(e)
	return e
}