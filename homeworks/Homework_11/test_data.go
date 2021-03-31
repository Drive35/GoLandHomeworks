package main

import (
	"GoLandHomeworks/homeworks/Homework_11/models"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var cl1, cl2, cl3 models.Client
	cl1 = models.Client{
		Id:       0001,
		FullName: "Auganbayev Madiyar",
		Balance:  10000,
	}
	cl2 = models.Client{
		Id:       0002,
		FullName: "Tlegen Daniyar",
		Balance:  15000,
	}
	cl3 = models.Client{
		Id:       0003,
		FullName: "Akhmetov Marat",
		Balance:  20000,
	}
	clients := []models.Client{cl1, cl2, cl3}
	fmt.Println(clients)
	clientsJson, err := json.Marshal(clients)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("clients.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(clientsJson)
	if err != nil {
		panic(err)
	}
}
