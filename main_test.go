package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestProcessReceiptHandler(t *testing.T) {
    payload := `{
        "retailer": "Target",
        "purchaseDate": "2022-01-02",
        "purchaseTime": "13:13",
        "total": "1.25",
        "items": [
            {"shortDescription": "Pepsi - 12-oz", "price": "1.50"}
        ]
    }`

    req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(payload)))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(processReceiptHandler)
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", rr.Code)
    }

    var resp map[string]string
    err = json.Unmarshal(rr.Body.Bytes(), &resp)
    if err != nil {
        t.Fatalf("Failed to parse response: %v", err)
    }
    if _, exists := resp["id"]; !exists {
        t.Errorf("Response body missing 'id' field: %v", rr.Body.String())
    }
}

func TestGetPointsHandler(t *testing.T) {
    ReceiptStore = map[string]int{
        "test-id": 100,
    }

    req, err := http.NewRequest("GET", "/receipts/test-id/points", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")
    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", rr.Code)
    }

    var resp map[string]int
    json.Unmarshal(rr.Body.Bytes(), &resp)
    if resp["points"] != 100 {
        t.Errorf("Expected points 100, got %d", resp["points"])
    }
}

func TestInvalidReceiptPayload(t *testing.T) {
	payload := `{}`

	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(processReceiptHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rr.Code)
	}
}
