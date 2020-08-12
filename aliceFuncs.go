package main

import (
	"strconv"
	"time"

	"github.com/azzzak/alice"
)

func checkfunc(req alice.Request, resp alice.Response) *alice.Response {
	if req.IsNewSession() {
		return resp.Text(firstMessage)
	} else if itemExists(helpQuestions, req.Command()) {
		return helpFunction(resp)
	} else if itemExists(abilityQuestions, req.Command()) {
		return showPossibilities(resp)
	} else if itemExists(controllCommands, req.Command()) {
		return resp.Text("Выполняю")
	} else if itemExists(stopCommands, req.Command()) {
		return resp.Text("Конец потока").EndSession()
	} else if itemExists(presentationQuestions, req.Command()) {
		showPresentation(1, req)
		text, tts := returnResponce(1)
		return resp.Text(text).TTS(tts)
	} else if itemExists(timeQuestions, req.Command()) {
		return printCurrentTime(resp)
	} else if req.Command() == "ня" {
		go func() *alice.Response {
			return resp.Text("test1")
		}()
		go func() *alice.Response {
			return resp.Text("test2")
		}()
		go func() *alice.Response {
			return resp.Text("test3")
		}()
		return resp.Text("test0")
	} else {
		return resp.Text("Не поняла вопроса, переформулируйте его, и повторите снова")
	}
}

func itemExists(array []string, item interface{}) bool {
	for _, value := range array {
		if value == item {
			return true
		}
	}
	return false
}

func searchIn(array []*alice.Request, item string) *alice.Request {
	for _, value := range array {
		if value.UserID() == item {
			return value
		}
	}
	return nil
}

func helpFunction(response alice.Response) *alice.Response {
	return response.Text(firstMessage)
}

func showPossibilities(response alice.Response) *alice.Response {
	return response.Text(possibilities)
}

func printCurrentTime(response alice.Response) *alice.Response {
	hour := time.Now().Hour()
	min := time.Now().Minute()
	year := time.Now().Year()
	mon := time.Now().Month()
	day := time.Now().Day()
	return response.Text("Сегодня " + strconv.Itoa(day) + ", " + mon.String() + ", " + strconv.Itoa(year) + ". Время: " + strconv.Itoa(hour) + " часов " + strconv.Itoa(min) + " минут ")
}
