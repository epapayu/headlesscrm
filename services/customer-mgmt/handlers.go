package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory mock database
var customers = []Customer{
	{
		ID:           "cust-001",
		Href:         "/customer/cust-001",
		Name:         "Budi Santoso",
		Status:       "Active",
		StatusReason: "KYC Verified",
		EngagedParty: PartyRef{ID: "party-001", Name: "Budi Santoso"},
		ContactMedium: []ContactMedium{
			{
				MediumType: "Phone",
				Preferred:  true,
				Characteristic: MediumCharacteristic{
					PhoneNumber: "+6281234567890",
				},
			},
		},
		Characteristic: []Characteristic{
			{Name: "NIK", Value: "3201123456789001"},
			{Name: "KK", Value: "3201987654321001"},
		},
	},
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	for _, cust := range customers {
		if cust.ID == id {
			c.JSON(http.StatusOK, cust)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

func ListCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	var newCust Customer
	if err := c.ShouldBindJSON(&newCust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pseudo-logic: Assign ID and mock KYC check
	newCust.ID = fmt.Sprintf("cust-%d", time.Now().Unix())
	newCust.Href = "/customer/" + newCust.ID
	newCust.Status = "Initial"
	newCust.StatusReason = "KYC Pending"

	// Add to mock DB
	customers = append(customers, newCust)

	c.JSON(http.StatusCreated, newCust)
}

func PatchCustomer(c *gin.Context) {
	id := c.Param("id")
	var patchReq map[string]interface{}
	if err := c.ShouldBindJSON(&patchReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, cust := range customers {
		if cust.ID == id {
			// Naive patch implementation for demo
			if status, ok := patchReq["status"].(string); ok {
				customers[i].Status = status
			}
			if reason, ok := patchReq["statusReason"].(string); ok {
				customers[i].StatusReason = reason
			}
			c.JSON(http.StatusOK, customers[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}
