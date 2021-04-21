package main

import (
	"GoLandHomeworks/homeworks/Homework_17/config"
	"GoLandHomeworks/homeworks/Homework_17/users"
	"fmt"
	"log"
)

func main() {
	user1 := users.User{
		Username: "Zhomart Zhanuzakov",
		Password: "qwerty11",
	}
	user2 := users.User{
		Username: "Ivan Ivanov",
		Password: "asdfgh224?",
	}
	user3 := users.User{
		Username: "Peter Johnson",
		Password: "zxcvbn33",
	}
	user4 := users.User{
		Username: "Sam Craige",
		Password: "gfhjkm44",
	}

	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
	}
	//добавление пользователей
	usersStore, err := users.NewUsersStore(cf)
	if err != nil {
		log.Fatal(err)
	}
	err = usersStore.AddUser(user1)
	if err != nil {
		log.Fatal(err)
	}
	err = usersStore.AddUser(user2)
	if err != nil {
		log.Fatal(err)
	}
	err = usersStore.AddUser(user3)
	if err != nil {
		log.Fatal(err)
	}
	err = usersStore.AddUser(user4)
	if err != nil {
		log.Fatal(err)
	}
	//вытаскивание одного пользователя по ID
	mongoUser, err := usersStore.GetById(467)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mongoUser)
	//вытаскивание полного списка пользователей
	listOfUsers, err := usersStore.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(listOfUsers)
	//update одного пользователя с Id 667
	user2 := users.User{
		Id: 667,
		Username: "Ivan Ivanov",
		Password: "asdfgh225!",
	}
	err = usersStore.Update(user2)
	if err != nil {
		log.Fatal(err)
	}
	//Удаление пользователя с Id 183
	oneUser, err := usersStore.GetById(183)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oneUser)
	err = usersStore.Delete(183)
	if err != nil {
		log.Fatal(err)
	}
}
