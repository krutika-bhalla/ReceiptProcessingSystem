package main

import (
	"log"
	"net/http"
	"receipt-processor/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/receipts/process", handlers.ProcessReceipt)
	router.GET("/receipts/:id/points", handlers.GetPoints)

	log.Fatal(http.ListenAndServe(":8080", router))
} 