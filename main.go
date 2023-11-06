package main

import (
	"lokasani/drivers/mysql"
	routes "lokasani/router"
)

func main() {
	db := mysql.InitDB()
	e := routes.Route(db)

	e.Start(":8081")
}
