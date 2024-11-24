package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessReceiptHandler(t *testing.T) {
	payload := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-02",
		"purchaseTime": "13:13",
		"total": "1.25",
		"items": [
			{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
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
	json.Unmarshal(rr.Body.Bytes(), &resp)
	if _, exists := resp["id"]; !exists {
		t.Errorf("Response body missing 'id' field: %v", rr.Body.String())
	}
}

func TestGetPointsHandler(t *testing.T) {
	id := "test-id"
	points := 100
	ReceiptStore[id] = points

	req, err := http.NewRequest("GET", "/receipts/"+id+"/points", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPointsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var resp map[string]int
	json.Unmarshal(rr.Body.Bytes(), &resp)
	if resp["points"] != points {
		t.Errorf("Expected points %d, got %d", points, resp["points"])
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
