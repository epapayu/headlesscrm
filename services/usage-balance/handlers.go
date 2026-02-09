package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory mock database
var balances = []AccountBalance{
	{
		ID:     "bal-123",
		Href:   "/balance/bal-123",
		Status: "Active",
		Amount: Money{Value: 50000, Unit: "IDR"},
		ValidFor: TimePeriod{
			StartDateTime: time.Now(),
			EndDateTime:   time.Now().AddDate(0, 1, 0),
		},
		RelatedParty: []RelatedParty{
			{ID: "cust-001", Name: "Budi Santoso", Role: "Owner"},
		},
	},
}

func GetBalance(c *gin.Context) {
	id := c.Param("id")
	for _, b := range balances {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Balance not found"})
}

func ListBalances(c *gin.Context) {
	// Simple mock: return all
	c.JSON(http.StatusOK, balances)
}

func AdjustBalance(c *gin.Context) {
	var adj BalanceAdjustment
	if err := c.ShouldBindJSON(&adj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pseudo-logic: Find wallet and update
	walletID := adj.RelatedWallet.ID
	found := false
	for i, b := range balances {
		if b.ID == walletID {
			if adj.Type == "deduction" {
				balances[i].Amount.Value -= adj.Amount.Value
			} else {
				balances[i].Amount.Value += adj.Amount.Value
			}
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	adj.ID = "adj-" + time.Now().Format("20060102150405")
	c.JSON(http.StatusCreated, adj)
}
