package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/azzzak/alice"
)

type Presentation struct {
	Name      string   `json:name`
	Actions   []action `json:actions`
	Timestamp string   `json:timestamp`
}

type action struct {
	Type string `json:type`
	Args string `json:args`
	Time int    `json:time`
}

var (
	assetsFolder string = "assets"
	preffix      string = "presentation"
)

func showPresentation(id int, req alice.Request) {
	workingDir := assetsFolder + "/" + preffix + "_" + strconv.Itoa(id)
	sessionDir := "Sessions/" + strconv.Itoa(id)
	data := openPresentation(id)
	for _, currentAction := range data.Actions {
		if currentAction.Type == "read" {
			timer := time.AfterFunc(time.Second, func() {
				showText(workingDir+"/text.txt", currentAction.Args)
			})
			defer timer.Stop()
		} else if currentAction.Type == "showImage" {
			timer := time.AfterFunc(time.Second, func() {
				showImage(workingDir+"images/"+currentAction.Args, sessionDir+"/img.png")
			})
			defer timer.Stop()
		}
		time.Sleep(time.Duration(currentAction.Time))
	}
}

func openPresentation(id int) Presentation {
	var presentation Presentation
	var workingDir string = assetsFolder + "/" + preffix + "_" + strconv.Itoa(id)
	presentationFile, err := os.Open(workingDir + "/data.json")
	errCheck(err)
	defer presentationFile.Close()
	presentationData, _ := ioutil.ReadAll(presentationFile)
	json.Unmarshal([]byte(presentationData), &presentation)
	return presentation
}

func returnResponce(id int) (string, string) {
	data := openPresentation(id)
	var responce string
	var tts string
	for _, currentAction := range data.Actions {
		if currentAction.Type == "read" {
			responce = responce + "\n" + currentAction.Args
			tts = tts + currentAction.Args + " sil <[" + strconv.Itoa(currentAction.Time) + "]> "
		}
	}
	return responce, tts
}

func showImage(assetsFile string, imgFile string) {
	data, err := ioutil.ReadFile(assetsFile)
	errCheck(err)
	err = ioutil.WriteFile(imgFile, data, 755)
}

func showText(textFile string, text string) {
	file, err := os.OpenFile(textFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 757)
	errCheck(err)
	file.WriteString(text)
	errCheck(err)
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
