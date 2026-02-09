package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateIdentity(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: NIK and KK are required"})
		return
	}

	// Mock Logic:
	// - Valid if NIK starts with "3201" (West Java)
	// - Invalid otherwise
	if strings.HasPrefix(req.NIK, "3201") && len(req.NIK) == 16 {
		c.JSON(http.StatusOK, ValidationResponse{
			IsValid:  true,
			Message:  "Identity Verified",
			FullName: "Mock Citizen Name",
		})
	} else {
		c.JSON(http.StatusOK, ValidationResponse{
			IsValid: false,
			Message: "Identity Not Found in Dukcapil",
		})
	}
}
