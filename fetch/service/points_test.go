package service

import (
	"receipt-processor/models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Target Receipt",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28, // Breakdown:
			// 6 points - retailer name has 6 characters
			// 10 points - 5 items (2 pairs @ 5 points each)
			// 3 Points - "Emils Cheese Pizza" is 18 characters (multiple of 3)
			// 3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (multiple of 3)
			// 6 points - purchase day is odd
		},
		{
			name: "M&M Corner Market Receipt",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109, // Breakdown:
			// 14 points - retailer name has 14 alphanumeric characters (& is not alphanumeric)
			// 50 points - total is a round dollar amount
			// 25 points - total is a multiple of 0.25
			// 10 points - 4 items (2 pairs @ 5 points each)
			// 10 points - 14:33 is between 2:00pm and 4:00pm
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatePoints(tt.receipt)
			if got != tt.expected {
				t.Errorf("CalculatePoints() = %v, want %v\nReceipt: %+v", got, tt.expected, tt.receipt)
			}
		})
	}
}

// TestCountAlphanumeric specifically tests the alphanumeric counting function
func TestCountAlphanumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Target",
			input:    "Target",
			expected: 6,
		},
		{
			name:     "M&M Corner Market",
			input:    "M&M Corner Market",
			expected: 14, // M, M, C, o, r, n, e, r, M, a, r, k, e, t
		},
		{
			name:     "Special Characters",
			input:    "A&B-C",
			expected: 3, // only A, B, C are counted
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countAlphanumeric(tt.input)
			if got != tt.expected {
				t.Errorf("countAlphanumeric(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
} 