package main

import (
	"encoding/json"
	"github.com/ZhomartZhan/GoLandHomeworks/homeworks/Homework_11/models"
	"io/ioutil"
	"os"
)

func main() {
	readdata, err := ioutil.ReadFile("clients.json")
	if err != nil {
		panic(err)
	}
	var clients2 []models.Client
	err = json.Unmarshal(readdata, &clients2)
	if err != nil {
		panic(err)
	}
	for i, _ := range clients2 {
		b := 1000
		clients2[i].DecreaseBalance(b)
	}
	clientsJson2, err := json.Marshal(clients2)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("clients.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(clientsJson2)
	if err != nil {
		panic(err)
	}
}
