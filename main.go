package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
	Points       int    `json:"points"`
}

var receiptData = make(map[string]Receipt)

func getPoints(context *gin.Context) {

	id := context.Param("id")

	// Retrieve the corresponding receipt from the map based on the ID
	receipt, ok := receiptData[id]

	if !ok {
		context.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found. ID not existed."})
		return
	}

	calculatePoints(&receipt)

	// Return the points in the response
	context.JSON(http.StatusOK, gin.H{"points": receipt.Points})

}

func processReceipt(context *gin.Context) {

	var receipt Receipt

	// Bind received JSON data to the Receipt struct

	if err := context.ShouldBindJSON(&receipt); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate random ID for the receipt

	id := uuid.New().String()
	receipt.ID = id

	receiptData[id] = receipt

	// Return the generated ID in the response
	context.JSON(http.StatusCreated, gin.H{"id": id})

}

func calculatePoints(receipt *Receipt) {

	var points = 0

	// One point for every alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == math.Floor(total) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += len(receipt.Items) / 2 * 5

	// Multiply the price by 0.2 and round up to the nearest integer if the trimmed length of the item description is a multiple of 3
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	timeLayout := "15:04"
	purchaseTime, err := time.Parse(timeLayout, receipt.PurchaseTime)
	startTime, _ := time.Parse(timeLayout, "14:00")
	endTime, _ := time.Parse(timeLayout, "16:00")
	if err == nil && purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		points += 10
	}

	receipt.Points = points
}

func countAlphanumeric(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func main() {
	router := gin.Default()

	router.GET("/receipts/:id/points", getPoints)
	router.POST("/receipts/process", processReceipt)

	// Start
	router.Run("localhost:3000")
}
