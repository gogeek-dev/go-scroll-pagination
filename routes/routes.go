package routes

import (
	"database/sql"
	"go_product_scroll/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*/*.html")

	r.Static("/assets", "./assets")

	r.GET("/", controller.Products)

	return r
}
