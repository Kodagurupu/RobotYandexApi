package main

import (
	"fmt"

	"github.com/azzzak/alice"
)

func itemExists(array []string, item interface{}) bool {
	fmt.Printf(array[0])
	for value := range array {
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
