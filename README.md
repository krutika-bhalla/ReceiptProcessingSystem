# Receipt Processing System

A web service that processes receipts and calculates reward points based on specific rules. This project is built in response to the [Fetch Rewards Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge).

## Features

- Process receipts and generate unique IDs
- Calculate points based on receipt data
- RESTful API endpoints
- In-memory data storage
- Points calculation based on multiple rules

## API Endpoints

### 1. Process Receipt
- **Endpoint**: `/receipts/process`
- **Method**: POST
- **Description**: Processes a receipt and returns a unique ID
- **Request Body**: JSON receipt data
- **Response**: JSON containing receipt ID

Example Request:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
  ],
  "total": "6.49"
}
```

Example Response:
```json
{
  "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

### 2. Get Points
- **Endpoint**: `/receipts/{id}/points`
- **Method**: GET
- **Description**: Retrieves points awarded for a specific receipt
- **Response**: JSON containing points awarded

Example Response:
```json
{
  "points": 32
}
```

## Points Calculation Rules

Points are awarded based on the following rules:

1. One point for every alphanumeric character in the retailer name
2. 50 points if the total is a round dollar amount with no cents
3. 25 points if the total is a multiple of 0.25
4. 5 points for every two items on the receipt
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer
6. 5 points if the total is greater than 10.00
7. 6 points if the day in the purchase date is odd
8. 10 points if the time of purchase is after 2:00pm and before 4:00pm

## Setup and Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ReceiptProcessingSystem.git
cd ReceiptProcessingSystem
```

2. Install dependencies:
```bash

go mod download
```

3. Run the application:
```bash
go run main.go
```

The application will start running on `http://localhost:8080`

## Testing

To run the tests:
```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 