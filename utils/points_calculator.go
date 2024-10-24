package utils

import (
    "math"
    "receipt-processor/models"
    "strconv"
    "strings"
    "unicode"
)

func isAlphanumeric(ch rune) bool {
    return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func isRoundDollarAmount(total string) bool {
    price, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return false
    }
    return price == float64(int(price))
}

func isMultipleOfQuarter(total string) bool {
    price, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return false
    }
    return int(price*100)%25 == 0
}

func extractDay(purchaseDate string) int {
    dateParts := strings.Split(purchaseDate, "-")
    if len(dateParts) != 3 {
        return 0
    }
    day, err := strconv.Atoi(dateParts[2])
    if err != nil {
        return 0
    }
    return day
}

func isBetweenTwoAndFourPM(purchaseTime string) bool {
    timeParts := strings.Split(purchaseTime, ":")
    if len(timeParts) != 2 {
        return false
    }

    hour, err := strconv.Atoi(timeParts[0])
    if err != nil {
        return false
    }
    
    return hour >= 14 && hour < 16
}

func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // Rule 1: Points for retailer name length
    for _, ch := range receipt.Retailer {
        if isAlphanumeric(ch) {
            points++
        }
        
    }

    // Rule 2: Round dollar amount
    if isRoundDollarAmount(receipt.Total) {
        points += 50
        
    }

    // Rule 3: Multiple of 0.25
    if isMultipleOfQuarter(receipt.Total) {
        points += 25
        
    }

    // Rule 4: 5 points per every two items
    points += (len(receipt.Items) / 2) * 5
    

    // Rule 5: Item description length
    for _, item := range receipt.Items {
        descriptionLength := len(strings.TrimSpace(item.ShortDescription))
        if descriptionLength%3 == 0 {
            price, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(price * 0.2))
            
        }
    }

    // Rule 6: Odd day
    day := extractDay(receipt.PurchaseDate)
    if day%2 != 0 {
        points += 6
        
    }

    // Rule 7: Time between 2:00pm and 4:00pm
    if isBetweenTwoAndFourPM(receipt.PurchaseTime) {
        points += 10
        
    }

    return points
}
