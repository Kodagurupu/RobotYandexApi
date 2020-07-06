package main

import (
	"log"
	"net/http"

	"github.com/azzzak/alice"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/", serve).Methods("POST")
	updates := alice.ListenForWebhook("/")
	go http.ListenAndServeTLS(":3000", "server.crt", "server.key", router)
	http.ListenAndServe(":8080", router)
	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()

		log.Printf("User send: " + req.OriginalUtterance())
		requests = append(requests, req)

		if req.IsNewSession() {
			return resp.Text("Здравствуйте. " + firstMessage)
		} else if itemExists(helpQuestions, req.Command()) {
			return helpFunction(*resp)
		} else if itemExists(abilityQuestions, req.Command()) {
			return showPossibilities(*resp)
		} else if itemExists(commands, req.Command()) {
			return resp.Text("Выполняю")
		} else {
			return resp.Text("Не поняла вопроса, переформулируйте его, и повторите снова")
		}
	})
}
