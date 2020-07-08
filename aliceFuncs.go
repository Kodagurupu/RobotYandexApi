package main

import (
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
	hour, min, _ := time.Now().Clock()
	year, mon, day := time.Now().Date()
	return response.Text("Сегодня "+string(day)+","+string(mon), ","+string(year)+", Время: "+string(hour)+" часов "+string(min)+" минут ")
}
