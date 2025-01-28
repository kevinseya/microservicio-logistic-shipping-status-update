package controllers

import (
	"net/http"
	"shippment-asignment/database"
	"shippment-asignment/models"

	"github.com/gin-gonic/gin"
)

func UpdateShipmentStatus(c *gin.Context) {
	var input struct {
		State string `json:"state" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipmentID := c.Param("id")

	// Convert UUIID to Bin for search on database
	var shipment models.Shipment
	if err := database.DB.Where("shipment_id = UUID_TO_BIN(?)", shipmentID).First(&shipment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}

	shipment.State = input.State
	if err := database.DB.Save(&shipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Shipment status updated successfully",
		"shipment": gin.H{
			"id":    shipmentID,
			"state": shipment.State,
		},
	})
}
