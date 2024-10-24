package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "receipt-processor/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

    http.ListenAndServe(":8080", r)
}
