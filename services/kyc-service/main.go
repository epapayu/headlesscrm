package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/kyc/v1")
	{
		v1.POST("/validate", ValidateIdentity)
	}

	r.Run(":8082") // Listen on port 8082
}
