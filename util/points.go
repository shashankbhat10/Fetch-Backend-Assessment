package util

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/shashankbhat10/Fetch-Backend-Assessment/models"
)

func GetPoints(receipt models.Receipt) (int, error) {
	points := 0
	points += getPointsForRetailerName(receipt.Retailer)
	pointsFromAmount, err := getPointsForAmount(receipt.Total)
	if err != nil {
		return 0, err
	}
	pointsFromItems, err := getPointsForItems(receipt.Items)
	if err != nil {
		return 0, err
	}
	pointsFromDateTime, err := getPointsForDateTime(receipt.PurchaseDate, receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}

	points += pointsFromAmount + pointsFromItems + pointsFromDateTime

	return points, nil
}

func getPointsForRetailerName(retailer string) int {
	count := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func getPointsForAmount(price string) (int, error) {
	parsedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, errors.New("invalid price")
	}
	points := 0
	priceInto100 := (int64)(parsedPrice * 100)
	if priceInto100%100 == 0 {
		points += 50
	}
	if priceInto100%25 == 0 {
		points += 25
	}
	return points, nil
}

func getPointsForItems(items []models.ReceiptItem) (int, error) {
	points := 0
	points = 5 * (len(items) / 2)

	for _, item := range items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, errors.New("incorrect price")
			}
			multipliedPrice := priceFloat * 0.2
			points += (int)(math.Ceil(multipliedPrice))
		}
	}
	return points, nil
}

func getPointsForDateTime(datestr string, timestr string) (int, error) {
	hour, err := strconv.ParseInt(strings.Split(timestr, ":")[0], 10, 0)
	if err != nil {
		return 0, errors.New("invalid time format")
	}
	minute, err := strconv.ParseInt(strings.Split(timestr, ":")[0], 10, 0)
	if err != nil {
		return 0, errors.New("invalid time format")
	}

	points := 0
	if hour >= 14 && hour <= 16 {
		if minute == 0 && (hour == 14 || hour == 16) {
			points += 0
		} else {
			points += 10
		}
	}

	fmt.Println(datestr)
	receiptDate, err := time.Parse(time.DateOnly, datestr)
	if err != nil {
		return 0, errors.New("invalid receipt date")
	}
	if receiptDate.Day()%2 == 1 {
		points += 6
	}

	return points, nil
}
