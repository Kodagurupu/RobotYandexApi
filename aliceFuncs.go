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

func helpFunction(response alice.Response) *alice.Response {
	return response.Text(firstMessage)
}

func showPossibilities(response alice.Response) *alice.Response {
	return response.Text(possibilities)
}
