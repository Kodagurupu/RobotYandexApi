package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/azzzak/alice"
)

func fileExists(filename string) bool {

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func createDirectory(dirName string) bool {

	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {

		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(err)
		}

		return true
	}

	if src.Mode().IsRegular() {
		log.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	createDirectory("Sessions")
}

func main() {
	updates := alice.ListenForWebhook("/")
	go http.ListenAndServeTLS(":3000", "server.crt", "server.key", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()

		log.Printf("User send: " + req.OriginalUtterance())
		file, err := os.OpenFile("Sessions/"+req.UserID(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 777)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		marshaled, _ := json.Marshal(req)
		file.WriteString(string(marshaled))
		if req.IsNewSession() {
			return resp.Text("Здравствуйте. " + firstMessage)
		} else if itemExists(helpQuestions, req.Command()) {
			return helpFunction(*resp)
		} else if itemExists(abilityQuestions, req.Command()) {
			return showPossibilities(*resp)
		} else if itemExists(controllCommands, req.Command()) {
			return resp.Text("Выполняю")
		} else if itemExists(stopCommands, req.Command()) {
			return resp.Text("Конец потока").EndSession()
		} else if itemExists(presentationQuestions, req.Command()) {
			return resp.Text(presentationAnswer)
		} else {
			return resp.Text("Не поняла вопроса, переформулируйте его, и повторите снова")
		}
	})
}
