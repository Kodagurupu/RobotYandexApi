package main

import (
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
