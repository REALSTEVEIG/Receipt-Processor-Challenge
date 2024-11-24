# Receipt Processor Challenge

This repository contains the solution to the **Receipt Processor Challenge** provided by Fetch Rewards. The application processes receipts and calculates points based on predefined rules.

## Features

- RESTful API built with Go.
- Two endpoints:
  - **POST `/receipts/process`**: Processes a receipt and returns a unique receipt ID.
  - **GET `/receipts/{id}/points`**: Retrieves the points awarded for a specific receipt.
- In-memory storage for simplicity (no database required).
- Fully Dockerized for ease of setup and deployment.

---

## How to Run the Application

### 1. Run Locally (Go Installed)

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/receipt-processor-challenge.git
   cd receipt-processor-challenge

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

4. The application will run on `http://localhost:8080`.

---

### 2. Run Using Docker

1. Build the Docker image:
   ```bash
   docker build -t receipt-processor .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 receipt-processor
   ```

3. The application will run on `http://localhost:8080`.

---

## API Endpoints

### 1. POST `/receipts/process`

**Description**: Accepts a receipt JSON payload, calculates points, and returns a unique receipt ID.

**Example Request**:
```json
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
```

**Example Response**:
```json
{
    "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

### 2. GET `/receipts/{id}/points`

**Description**: Retrieves the points awarded for a specific receipt.

**Example Request**:
```bash
curl http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```

**Example Response**:
```json
{
    "points": 32
}
```

---

## Testing

### 1. Using Postman
- Import the provided Postman collection (if included).
- Test the endpoints with the example JSON payloads.

### 2. Using curl
```bash
# Process a receipt
curl -X POST -H "Content-Type: application/json" -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}]
}' http://localhost:8080/receipts/process

# Get points for a receipt ID
curl http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```

---

## Testing Suite

Unit tests are included in the `*_test.go` files. Run tests using:
```bash
go test ./...
```

---

## Notes

- The application is designed to be stateless. Receipt data is stored in memory and will be cleared upon restarting the application.
- Ensure Docker or Go is installed before running the application.

---

## Assumptions

- Receipt validation follows the rules provided in the challenge.
- Time is handled in the `UTC` timezone for consistency.
```
