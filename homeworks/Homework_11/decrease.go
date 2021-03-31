package main

import (
	"GoLandHomeworks/homeworks/Homework_11/models"
	"encoding/json"
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
	for _, v := range clients2 {
		b := 1000
		v.DecreaseBalance(b)
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
