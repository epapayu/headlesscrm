package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/serviceOrderManagement/v4")
	{
		v4.GET("/serviceOrder", ListServiceOrders)
		v4.POST("/serviceOrder", CreateServiceOrder)
		v4.GET("/serviceOrder/:id", GetServiceOrder)
	}

	r.Run(":8087") // Listen on port 8087
}
