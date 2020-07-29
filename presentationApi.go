package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Presentation struct {
	Name      string   `json:name`
	Actions   []action `json:actions`
	Timestamp string   `json:timestamp`
}

type action struct {
	Type string `json:type`
	Args string `json:args`
}

var (
	assetsFolder string = "assets"
	preffix      string = "presentation"
)

func showPresentation(id int) {
	data := openPresentation(id)
	for _, currentAction := range data.Actions {
		if currentAction.Type == "read" {

		} else if currentAction.Type == "showImage" {

		}
	}
}

func openPresentation(id int) Presentation {
	var presentation Presentation
	var workingDir string = assetsFolder + preffix + strconv.Itoa(id)
	presentationFile, err := os.Open(workingDir + "data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer presentationFile.Close()
	presentationData, _ := ioutil.ReadAll(presentationFile)
	json.Unmarshal([]byte(presentationData), &presentation)
	return presentation
}

func returnResponce(id int) string {
	data := openPresentation(id)
	var responce string
	for _, currentAction := range data.Actions {
		if currentAction.Type == "read" {
			responce = responce + currentAction.Args + "sil <[1000]>"
		}
	}
	return responce
}
