package handlers

import (
    "encoding/json"
    "net/http"
    "receipt-processor/models"
    "receipt-processor/utils"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

var receipts = make(map[string]int)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    points := utils.CalculatePoints(receipt)

    receiptID := uuid.New().String()
    receipts[receiptID] = points

    response := map[string]string{"id": receiptID}
    json.NewEncoder(w).Encode(response)
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    receiptID := vars["id"]

    points, exists := receipts[receiptID]
    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    response := models.PointsResponse{Points: points}
    json.NewEncoder(w).Encode(response)
}
