package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Mock Data
var tickets = []TroubleTicket{
	{
		ID:           "TKT-001",
		Href:         "/troubleTicket/TKT-001",
		Description:  "Cannot access 5G network in Jakarta Selatan",
		Severity:     "Major",
		Type:         "NetworkFault",
		CreationDate: time.Now().Add(-24 * time.Hour),
		Status:       "InProgress",
		RelatedParty: []RelatedParty{{ID: "C-001", Role: "Customer", Name: "Budi Santoso"}},
		Note: []Note{
			{Date: time.Now().Add(-2 * time.Hour), Author: "System", Text: "Ticket assigned to Network Ops"},
		},
	},
}

func ListTroubleTickets(c *gin.Context) {
	c.JSON(http.StatusOK, tickets)
}

func RetrieveTroubleTicket(c *gin.Context) {
	id := c.Param("id")
	for _, t := range tickets {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
}

func CreateTroubleTicket(c *gin.Context) {
	var newTicket TroubleTicket
	if err := c.ShouldBindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTicket.ID = fmt.Sprintf("TKT-%d", time.Now().Unix())
	newTicket.Href = fmt.Sprintf("/troubleTicket/%s", newTicket.ID)
	newTicket.CreationDate = time.Now()
	newTicket.Status = "Submitted"
	tickets = append(tickets, newTicket)
	c.JSON(http.StatusCreated, newTicket)
}

func PatchTroubleTicket(c *gin.Context) {
	id := c.Param("id")
	var updateData struct {
		Status             string `json:"status"`
		StatusChangeReason string `json:"statusChangeReason"`
		Note               []Note `json:"note"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tickets {
		if t.ID == id {
			if updateData.Status != "" {
				tickets[i].Status = updateData.Status
			}
			if updateData.StatusChangeReason != "" {
				tickets[i].StatusChangeReason = updateData.StatusChangeReason
			}
			if len(updateData.Note) > 0 {
				tickets[i].Note = append(tickets[i].Note, updateData.Note...)
			}
			c.JSON(http.StatusOK, tickets[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
}
