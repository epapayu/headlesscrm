package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/productInventoryManagement/v4")
	{
		v4.GET("/product", ListProducts)
		v4.GET("/product/:id", GetProduct)
	}

	r.Run(":8086") // Listen on port 8086
}
