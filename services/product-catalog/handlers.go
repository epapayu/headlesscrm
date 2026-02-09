package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mock Data
var offerings = []ProductOffering{
	{
		ID:          "offer-001",
		Href:        "/productOffering/offer-001",
		Name:        "50GB Freedom Data",
		Description: "50GB Data for 30 Days",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{
			{
				ID:        "price-001",
				Name:      "Standard Price",
				PriceType: "oneTime",
				Price:     Money{Value: 100000, Unit: "IDR"},
			},
		},
		Category: []CategoryRef{
			{ID: "cat-data", Name: "Data Packages"},
		},
	},
	{
		ID:          "offer-002",
		Href:        "/productOffering/offer-002",
		Name:        "Unlimited WhatsApp",
		Description: "Unlimited WhatsApp access for 7 days",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{
			{
				ID:        "price-002",
				Name:      "Weekly Price",
				PriceType: "oneTime",
				Price:     Money{Value: 15000, Unit: "IDR"},
			},
		},
		Category: []CategoryRef{
			{ID: "cat-social", Name: "Social Media Add-ons"},
		},
	},
}

func ListProductOfferings(c *gin.Context) {
	c.JSON(http.StatusOK, offerings)
}

func GetProductOffering(c *gin.Context) {
	id := c.Param("id")
	for _, off := range offerings {
		if off.ID == id {
			c.JSON(http.StatusOK, off)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product Offering not found"})
}

func CreateProductOffering(c *gin.Context) {
	var newOff ProductOffering
	if err := c.ShouldBindJSON(&newOff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Simplified mock creation
	newOff.ID = "offer-new"
	newOff.Status = "Active"
	offerings = append(offerings, newOff)
	c.JSON(http.StatusCreated, newOff)
}
