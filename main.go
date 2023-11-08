package main

import (
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"
)

func main() {
	db := mysql.InitDB()
	e := routes.Route(db)

	e.Start(":8081")
}
