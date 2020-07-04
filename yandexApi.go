package main

import (
	"net/http"

	"github.com/azzzak/alice"
)

func main() {
	updates := alice.ListenForWebhook("/api/yandex")
	go http.ListenAndServeTLS(":3000", "server.crt", "server.crt", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("Здравствуйте")
		}
		return resp.Text(req.OriginalUtterance())
	})
}
