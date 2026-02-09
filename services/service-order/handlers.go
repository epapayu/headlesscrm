package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var serviceOrders = []ServiceOrder{}

func CreateServiceOrder(c *gin.Context) {
	var newOrder ServiceOrder
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pseudo-logic: Assign ID and initial state
	newOrder.ID = fmt.Sprintf("srv-ord-%d", time.Now().Unix())
	newOrder.Href = "/serviceOrder/" + newOrder.ID
	newOrder.State = "Acknowledged"
	newOrder.OrderDate = time.Now()

	// Stub: Network Provisioning
	go provisionService(newOrder.ID)

	serviceOrders = append(serviceOrders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}

func GetServiceOrder(c *gin.Context) {
	id := c.Param("id")
	for _, ord := range serviceOrders {
		if ord.ID == id {
			c.JSON(http.StatusOK, ord)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Service Order not found"})
}

func ListServiceOrders(c *gin.Context) {
	c.JSON(http.StatusOK, serviceOrders)
}

// Mock Provisioning Logic
func provisionService(orderID string) {
	fmt.Printf("Provisioning Service Order %s...\n", orderID)
	time.Sleep(2 * time.Second)
	// Update state to Completed
	for i, ord := range serviceOrders {
		if ord.ID == orderID {
			serviceOrders[i].State = "Completed"
			fmt.Printf("Service Order %s Provisioned (Completed)\n", orderID)
			break
		}
	}
}
