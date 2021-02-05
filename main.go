package main

import (
	mysqldb "go_product_scroll/connection"
	"go_product_scroll/routes"
	"log"
)

func main() {

	db := mysqldb.SetupDB()

	r := routes.SetupRoutes(db)

	log.Println("listening on http://localhost:8030")

	r.Run(":8030")

}
