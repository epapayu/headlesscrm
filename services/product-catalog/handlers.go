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

	// Roaming
	{
		ID:          "offer-roam-asia",
		Name:        "Roaming Asia",
		Description: "7 Days 5GB Asia Roaming",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-roam-asia", Name: "Roaming Asia Price", PriceType: "oneTime", Price: Money{Value: 150000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-roam", Name: "Roaming"}},
	},
	{
		ID:          "offer-roam-eu",
		Name:        "Roaming Europe",
		Description: "14 Days 10GB Europe Roaming",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-roam-eu", Name: "Roaming EU Price", PriceType: "oneTime", Price: Money{Value: 350000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-roam", Name: "Roaming"}},
	},
	{
		ID:          "offer-roam-us",
		Name:        "Roaming US",
		Description: "14 Days 10GB US Roaming",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-roam-us", Name: "Roaming US Price", PriceType: "oneTime", Price: Money{Value: 400000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-roam", Name: "Roaming"}},
	},
	// Streaming
	{
		ID:          "offer-netflix",
		Name:        "Netflix Mobile",
		Description: "Netflix Mobile Plan 1 Month",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-netflix", Name: "Netflix Price", PriceType: "recurring", Price: Money{Value: 54000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-ent", Name: "Entertainment"}},
	},
	{
		ID:          "offer-spotify",
		Name:        "Spotify Premium",
		Description: "Spotify Premium Individual 1 Month",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-spotify", Name: "Spotify Price", PriceType: "recurring", Price: Money{Value: 55000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-ent", Name: "Entertainment"}},
	},
	{
		ID:          "offer-disney",
		Name:        "Disney+ Hotstar",
		Description: "Disney+ Hotstar Basic 1 Month",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-disney", Name: "Disney Price", PriceType: "recurring", Price: Money{Value: 39000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-ent", Name: "Entertainment"}},
	},
	// Gaming
	{
		ID:          "offer-games-max",
		Name:        "GamesMAX Power",
		Description: "3GB Data + 15GB Game Quota",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-games", Name: "GamesMAX Price", PriceType: "oneTime", Price: Money{Value: 75000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-game", Name: "Gaming"}},
	},
	// Postpaid
	{
		ID:          "offer-post-lite",
		Name:        "Postpaid Lite",
		Description: "15GB Data + 100 Mins Voice",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-post-lite", Name: "Lite Plan", PriceType: "recurring", Price: Money{Value: 100000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-post", Name: "Postpaid"}},
	},
	{
		ID:          "offer-post-pro",
		Name:        "Postpaid Pro",
		Description: "50GB Data + 300 Mins Voice",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-post-pro", Name: "Pro Plan", PriceType: "recurring", Price: Money{Value: 250000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-post", Name: "Postpaid"}},
	},
	{
		ID:          "offer-post-ultra",
		Name:        "Postpaid Ultra",
		Description: "Unlimited Data + Unlimited Voice",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-post-ultra", Name: "Ultra Plan", PriceType: "recurring", Price: Money{Value: 500000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-post", Name: "Postpaid"}},
	},
	// AI Bundles
	{
		ID:          "offer-ai-plus",
		Name:        "Google AI Plus",
		Description: "Gemini Advanced Access",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-ai-plus", Name: "AI Plus Price", PriceType: "recurring", Price: Money{Value: 300000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-ai", Name: "AI Bundles"}},
	},
	{
		ID:          "offer-ai-pro",
		Name:        "Google AI Pro",
		Description: "Gemini 1.5 Pro API Access (1M Tokens)",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-ai-pro", Name: "AI Pro Price", PriceType: "recurring", Price: Money{Value: 750000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-ai", Name: "AI Bundles"}},
	},
	{
		ID:          "offer-ai-prem",
		Name:        "Google AI Premium",
		Description: "Enterprise AI Suite + 5TB Cloud Storage",
		Status:      "Active",
		ProductOfferingPrice: []ProductOfferingPrice{{ID: "p-ai-prem", Name: "AI Premium Price", PriceType: "recurring", Price: Money{Value: 1500000, Unit: "IDR"}}},
		Category: []CategoryRef{{ID: "cat-ai", Name: "AI Bundles"}},
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
