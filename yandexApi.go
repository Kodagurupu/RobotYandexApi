package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azzzak/alice"
)

func main() {
	updates := alice.ListenForWebhook("/")
	go http.ListenAndServeTLS(":3000", "server.crt", "server.key", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("Здравствуйте")
		}
		fmt.Printf("User send" + req.OriginalUtterance())
		log.Printf("User send" + req.OriginalUtterance())
		return resp.Text(req.OriginalUtterance())
	})
}
