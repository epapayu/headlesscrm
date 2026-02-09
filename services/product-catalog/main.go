package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/productCatalogManagement/v4")
	{
		v4.GET("/productOffering", ListProductOfferings)
		v4.GET("/productOffering/:id", GetProductOffering)
		v4.POST("/productOffering", CreateProductOffering)
	}

	r.Run(":8083") // Listen on port 8083
}
