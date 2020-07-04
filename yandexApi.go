package main

import (
	"net/http"

	"github.com/azzzak/alice"
)

func main() {
	updates := alice.ListenForWebhook("/")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("Здравствуйте")
		}
		return resp.Text(req.OriginalUtterance())
	})
}
