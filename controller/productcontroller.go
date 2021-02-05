package controller

import (
	"fmt"
	mysqldb "go_product_scroll/connection"
	"go_product_scroll/models"

	"github.com/gin-gonic/gin"
)

func Products(c *gin.Context) {

	db := mysqldb.SetupDB()

	ProductsRows, err := db.Query("select id,title,description,price,image_path from tbl_products")

	if err != nil {

		fmt.Println(err)
	}
	products := models.Products{}

	res := []models.Products{}

	for ProductsRows.Next() {
		var id int

		var title, description, image_path string

		var price float32

		err = ProductsRows.Scan(&id, &title, &description, &price, &image_path)

		if err != nil {
			fmt.Println(err)
		}

		products.ID = id

		products.Title = title

		products.Price = price

		products.Imagepath = image_path

		res = append(res, products)

	}

	c.HTML(200, "products.html", res)

}
