package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/azzzak/alice"
)

func itemExists(array []string, item interface{}) bool {
	for _, value := range array {
		if value == item {
			return true
		}
	}
	return false
}

func searchIn(array []*alice.Request, item string) *alice.Request {
	for _, value := range array {
		if value.UserID() == item {
			return value
		}
	}
	return nil
}

func helpFunction(response alice.Response) *alice.Response {
	return response.Text(firstMessage)
}

func showPossibilities(response alice.Response) *alice.Response {
	return response.Text(possibilities)
}

func printCurrentTime(response alice.Response) *alice.Response {
	hour := time.Now().Hour()
	min := time.Now().Minute()
	year := time.Now().Year()
	mon := time.Now().Month()
	day := time.Now().Day()
	var timestamp = "Сегодня " + strconv.Itoa(day) + "," + mon.String() + "," + strconv.Itoa(year) + ", Время: " + strconv.Itoa(hour) + " часов " + strconv.Itoa(min) + " минут "
	fmt.Printf(timestamp, time.Now().String())
	return response.Text(timestamp)
}
