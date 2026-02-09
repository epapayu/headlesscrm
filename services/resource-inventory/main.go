package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/resourceInventoryManagement/v4")
	{
		v4.GET("/resource", ListResources)
		v4.GET("/resource/:id", GetResource)
	}

	r.Run(":8088") // Listen on port 8088
}
