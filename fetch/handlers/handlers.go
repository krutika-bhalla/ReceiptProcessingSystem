package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"receipt-processor/models"
	"receipt-processor/service"
)

var receipts = make(map[string]models.Receipt)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt

	c.JSON(http.StatusOK, models.ProcessResponse{ID: id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	receipt, exists := receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	points := service.CalculatePoints(receipt)
	c.JSON(http.StatusOK, models.PointsResponse{Points: points})
} 