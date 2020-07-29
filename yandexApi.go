package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/azzzak/alice"
)

func main() {
	updates := alice.ListenForWebhook("/")
	go http.ListenAndServeTLS(":3000", "server.crt", "server.key", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if !req.IsNewSession() {
			log.Printf("User send: " + req.OriginalUtterance())
		} else {
			configureUser(req.UserID())
		}
		file, err := os.OpenFile("Sessions/"+req.UserID()+"/responce.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 757)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
			return checkfunc(*req, *resp)
		} else {
			marshaled, _ := json.Marshal(req)
			file.WriteString(string(marshaled))
			return checkfunc(*req, *resp)
		}
	})
}

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

func configureUser(userID string) {
	if !fileExists("Sessions/" + userID) {
		createDirectory("Sessions/" + userID)
	}
}
