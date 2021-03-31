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
	var clients3 []models.Client
	err = json.Unmarshal(readdata, &clients3)
	if err != nil {
		panic(err)
	}
	for _, v := range clients3 {
		a := 1000
		v.IncreaseBalance(a)
	}
	clientsJson3, err := json.Marshal(clients3)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("clients.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(clientsJson3)
	if err != nil {
		panic(err)
	}
}
