package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var orders = []ProductOrder{}

func CreateProductOrder(c *gin.Context) {
	var newOrder ProductOrder
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pseudo-logic: Assign ID and initial state
	newOrder.ID = fmt.Sprintf("ord-%d", time.Now().Unix())
	newOrder.Href = "/productOrder/" + newOrder.ID
	newOrder.State = "Acknowledged"
	newOrder.OrderDate = time.Now()

	// Stub: Decomposition / Orchestration Trigger
	go orchestrateOrder(newOrder.ID)

	orders = append(orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}

func GetProductOrder(c *gin.Context) {
	id := c.Param("id")
	for _, ord := range orders {
		if ord.ID == id {
			c.JSON(http.StatusOK, ord)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func ListProductOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

// Mock orchestration logic
func orchestrateOrder(orderID string) {
	fmt.Printf("Orchestrating Order %s...\n", orderID)
	// Simulate processing time
	time.Sleep(2 * time.Second)
	// Update state to InProgress
	for i, ord := range orders {
		if ord.ID == orderID {
			orders[i].State = "InProgress"
			fmt.Printf("Order %s status updated to InProgress\n", orderID)
			break
		}
	}
}
