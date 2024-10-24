package tests

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "receipt-processor/handlers"
)

func TestProcessReceipt(t *testing.T) {
    receiptJSON := `{
        "retailer": "Target",
        "purchaseDate": "2022-01-02",
        "purchaseTime": "13:13",
        "total": "1.25",
        "items": [{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}]
    }`

    req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.ProcessReceiptHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}
