package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v4 := r.Group("/tmf-api/troubleTicketManagement/v4")
	{
		v4.GET("/troubleTicket", ListTroubleTickets)
		v4.GET("/troubleTicket/:id", RetrieveTroubleTicket)
		v4.POST("/troubleTicket", CreateTroubleTicket)
		v4.PATCH("/troubleTicket/:id", PatchTroubleTicket)
	}
	r.Run(":8089") // Port 8089 for Trouble Ticket Service
}
