package main

import (
	"GoLandHomeworks/homeworks/Homework_11/models"
	"encoding/json"
	"io/ioutil"
)

func main() {
	readdata, err := ioutil.ReadFile("clients.json")
	if err != nil {
		panic(err)
	}
	var clients []models.Client
	err = json.Unmarshal(readdata, &clients)
	if err != nil {
		panic(err)
	}
	for _, v := range clients {
		v.PrintAll()
	}
}