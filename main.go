package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "receipt-processor/docs"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        string  `json:"total"`
	Items        []Item  `json:"items"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var ReceiptStore = map[string]int{}

// @Summary Process a receipt
// @Description Submits a receipt for processing and returns a unique receipt ID.
// @Tags receipts
// @Accept  json
// @Produce  json
// @Param   receipt body Receipt true "Receipt to process"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid receipt"
// @Router /receipts/process [post]
func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validateReceipt(receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	points := calculatePoints(receipt)
	id := uuid.New().String()
	ReceiptStore[id] = points

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Get receipt points
// @Description Retrieves the points awarded for a receipt using its ID.
// @Tags receipts
// @Accept  json
// @Produce  json
// @Param   id path string true "Receipt ID"
// @Success 200 {object} map[string]int
// @Failure 404 {string} string "Receipt not found"
// @Router /receipts/{id}/points [get]
func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := ReceiptStore[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func validateReceipt(receipt Receipt) error {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" || len(receipt.Items) == 0 {
		return errors.New("missing required fields in receipt")
	}

	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return errors.New("invalid purchaseDate format")
	}
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return errors.New("invalid purchaseTime format")
	}

	if _, err := strconv.ParseFloat(receipt.Total, 64); err != nil {
		return errors.New("invalid total format")
	}

	for _, item := range receipt.Items {
		if item.ShortDescription == "" || item.Price == "" {
			return errors.New("invalid item in receipt")
		}
		if _, err := strconv.ParseFloat(item.Price, 64); err != nil {
			return errors.New("invalid item price format")
		}
	}

	return nil
}

func calculatePoints(receipt Receipt) int {
	points := 0

	alnum := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alnum.FindAllString(receipt.Retailer, -1))

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 == 1 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", processReceiptHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
