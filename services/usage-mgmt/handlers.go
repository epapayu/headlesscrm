package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var usageRecords = []Usage{}

func CreateUsage(c *gin.Context) {
	var newUsage Usage
	if err := c.ShouldBindJSON(&newUsage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pseudo-logic: Assign ID
	newUsage.ID = fmt.Sprintf("usage-%d", time.Now().Unix())
	newUsage.Href = "/usage/" + newUsage.ID
	newUsage.Date = time.Now()
	newUsage.Status = "Received"

	// Mock Rating Logic
	rateUsage(&newUsage)

	usageRecords = append(usageRecords, newUsage)
	c.JSON(http.StatusCreated, newUsage)
}

func GetUsage(c *gin.Context) {
	id := c.Param("id")
	for _, u := range usageRecords {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usage record not found"})
}

func ListUsage(c *gin.Context) {
	c.JSON(http.StatusOK, usageRecords)
}

func rateUsage(u *Usage) {
	// Simple mock rating
	// Data: 10 IDR per MB
	// Voice: 100 IDR per Minute
	// SMS: 50 IDR per SMS
	var amount float64 = 0

	for _, char := range u.UsageCharacteristic {
		if u.Type == "Data" && char.Name == "volume_mb" {
			vol, _ := strconv.ParseFloat(char.Value, 64)
			amount = vol * 10
		} else if u.Type == "Voice" && char.Name == "duration_sec" {
			dur, _ := strconv.ParseFloat(char.Value, 64)
			amount = (dur / 60) * 100
		}
	}

	if amount > 0 {
		u.Status = "Rated"
		u.RatedProductUsage = []RatedProductUsage{
			{
				RatingAmountType: "Total",
				TaxIncludedRatingAmount: Money{
					Value: amount,
					Unit:  "IDR",
				},
			},
		}
	}
}
