package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v4 := r.Group("/tmf-api/accountBalanceManagement/v4")
	{
		v4.GET("/balance", ListBalances)
		v4.GET("/balance/:id", GetBalance)
		v4.POST("/balanceAdjustment", AdjustBalance)
	}

	r.Run(":8080") // Listen on port 8080
}
