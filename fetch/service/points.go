package service

import (
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	roundedCents := int(math.Round(total * 100))
	if roundedCents%100 == 0 {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if roundedCents%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	pairs := len(receipt.Items) / 2 // integer division
	points += pairs * 5

	// Rule 5: Points for items with description length multiple of 3
	for _, item := range receipt.Items {
		shortDesc := strings.TrimSpace(item.ShortDescription)
		if len(shortDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	hour := purchaseTime.Hour()
	if hour >= 14 && hour < 16 {
		points += 10
	}

	return points
}

func countAlphanumeric(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1))
}
