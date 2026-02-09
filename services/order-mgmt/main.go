package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/productOrderManagement/v4")
	{
		v4.GET("/productOrder", ListProductOrders)
		v4.POST("/productOrder", CreateProductOrder)
		v4.GET("/productOrder/:id", GetProductOrder)
	}

	r.Run(":8084") // Listen on port 8084
}
