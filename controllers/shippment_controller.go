package controllers

import (
	"log"
	"net/http"
	"shippment-asignment/database"
	"shippment-asignment/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	log.Println("Received Shipment ID:", shipmentID)

	uuidID, err := uuid.Parse(shipmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	log.Println("Parsed UUID:", uuidID)

	var shipment models.Shipment
	if err := database.DB.Where("id = ?", uuidID).First(&shipment).Error; err != nil {
		log.Println("Shipment not found with ID:", shipmentID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}

	shipment.Status = input.State
	if err := database.DB.Save(&shipment).Error; err != nil {
		log.Println("Error saving shipment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Shipment status updated successfully",
		"shipment": gin.H{
			"id":         shipment.ID.String(),
			"state":      shipment.Status,
			"created_at": shipment.CreatedAt,
			"updated_at": shipment.UpdatedAt,
		},
	})
}
