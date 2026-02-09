package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mock Resources
var resources = []Resource{
	{
		ID:             "res-msisdn-001",
		Href:           "/resource/res-msisdn-001",
		Name:           "628123456789",
		Category:       "MSISDN",
		Description:    "Premium Number",
		ResourceStatus: "InUse",
		ResourceCharacteristic: []Characteristic{
			{Name: "prefix", Value: "62812"},
		},
	},
	{
		ID:             "res-sim-001",
		Href:           "/resource/res-sim-001",
		Name:           "896200000000001",
		Category:       "SIM",
		Description:    "USIM Card 64K",
		ResourceStatus: "InUse",
		ResourceCharacteristic: []Characteristic{
			{Name: "imsi", Value: "510100000000001"},
		},
	},
	{
		ID:             "res-msisdn-002",
		Href:           "/resource/res-msisdn-002",
		Name:           "628129999999",
		Category:       "MSISDN",
		Description:    "Gold Number",
		ResourceStatus: "Available",
		ResourceCharacteristic: []Characteristic{
			{Name: "prefix", Value: "62812"},
		},
	},
}

func ListResources(c *gin.Context) {
	category := c.Query("category")
	if category != "" {
		filtered := []Resource{}
		for _, r := range resources {
			if r.Category == category {
				filtered = append(filtered, r)
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}
	c.JSON(http.StatusOK, resources)
}

func GetResource(c *gin.Context) {
	id := c.Param("id")
	for _, r := range resources {
		if r.ID == id {
			c.JSON(http.StatusOK, r)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
}
