package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/usageManagement/v4")
	{
		v4.GET("/usage", ListUsage)
		v4.POST("/usage", CreateUsage)
		v4.GET("/usage/:id", GetUsage)
	}

	r.Run(":8085") // Listen on port 8085
}
