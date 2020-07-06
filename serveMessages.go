package main

import (
	"encoding/json"
	"net/http"
)

type request struct {
	RequestedID string `json:"requestedID"`
}

func serve(w http.ResponseWriter, r *http.Request) {

	var (
		form     request
		clientID string
	)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&form)

	clientID = form.RequestedID
	var response = searchIn(requests, clientID)

	json.NewEncoder(w).Encode(response)
}
