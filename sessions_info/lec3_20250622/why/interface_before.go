package main

import "sort"

func LogToConsole(msg string) {
    fmt.Println(msg)
}

func LogToFile(msg string) {
    // open file, write msg, close file
}

func LogToAPI(msg string) {
    // call REST API with msg
}


func PrintReport(mode string, msg string) {
	if mode == "console" {
		LogToConsole(msg)
	} else if mode == "file" {
		LogToFile(msg)
	} else if mode == "api" {
		LogToAPI(msg)
	}
	return
}


// ----------------------------
// if u wanna add a new logger 
func LogToSlack(msg string) {
    // send to Slack API
}

func defaultLog(msg string) {
	LogToConsole(msg)
}

func PrintReport(mode string, msg string) {

	if mode == "console" {
		LogToConsole(msg)
	} else if mode == "file" {
		LogToFile(msg)
	} else if mode == "api" {
		LogToAPI(msg)
	} else if mode == "slack" {
		LogToSlack(msg)
	}
	else defaultLog(msg)
	sort.Sort()
	return
}	
