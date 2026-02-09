package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Mock Inventory
var products = []Product{
	{
		ID:          "prod-inv-001",
		Href:        "/product/prod-inv-001",
		Name:        "50GB Freedom Data",
		Description: "Targeted Data Plan",
		Status:      "Active",
		StartDate:   time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
		ProductOffering: ProductOfferingRef{
			ID:   "offer-001",
			Name: "50GB Freedom Data",
		},
		RelatedParty: []RelatedParty{
			{ID: "cust-001", Role: "Customer", Name: "Budi Santoso"},
		},
		ProductCharacteristic: []Characteristic{
			{Name: "data_quota", Value: "50GB"},
		},
	},
}

func ListProducts(c *gin.Context) {
	customerID := c.Query("relatedParty.id")
	if customerID != "" {
		filtered := []Product{}
		for _, p := range products {
			for _, rp := range p.RelatedParty {
				if rp.ID == customerID {
					filtered = append(filtered, p)
					break
				}
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
