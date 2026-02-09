package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/customerManagement/v4")
	{
		v4.GET("/customer", ListCustomers)
		v4.POST("/customer", CreateCustomer)
		v4.GET("/customer/:id", GetCustomer)
		v4.PATCH("/customer/:id", PatchCustomer)
	}

	r.Run(":8081") // Listen on port 8081 (distinct from 8080)
}
