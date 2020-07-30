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
	sessionDir := "Sessions/" + req.UserID()
	data := openPresentation(id)
	for index, currentAction := range data.Actions {
		timer := time.NewTimer(time.Second*time.Duration(index*2) + time.Millisecond*time.Duration(index*100))
		log.Println(index * 2 * 1000)
		if currentAction.Type == "read" {
			go func(sessionDir string, currentAction action) {
				<-timer.C
				log.Println(currentAction.Type, currentAction.Args)
				showText(sessionDir+"/text.txt", currentAction.Args)
			}(sessionDir, currentAction)
		} else if currentAction.Type == "showImage" {
			go func(workingDir string, currentAction action, sesssessionDir string) {
				<-timer.C
				log.Println(currentAction.Type, currentAction.Args)
				showImage(workingDir+"/images/"+currentAction.Args, sessionDir+"/img.png")
			}(workingDir, currentAction, sessionDir)
		}
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
	for index, currentAction := range data.Actions {
		var await int
		if index == 0 {
			await = 0
		} else if index+1 != len(data.Actions) {
			await = 2000
		} else {
			await = 0
		}
		if currentAction.Type == "read" {
			responce = responce + "\n" + currentAction.Args
			tts = tts + currentAction.Args + " sil <[" + strconv.Itoa(await) + "]> "
		} else {
			tts = tts + " sil <[" + strconv.Itoa(await) + "]> "
		}
	}
	return responce, tts
}

func showImage(assetsFile string, imgFile string) {
	data, err := ioutil.ReadFile(assetsFile)
	errCheck(err)
	if fileExists(imgFile) {
		os.Remove(imgFile)
	}
	err = ioutil.WriteFile(imgFile, data, 755)
	errCheck(err)
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
