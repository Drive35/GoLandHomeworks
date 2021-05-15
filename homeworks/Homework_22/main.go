package main

import (
	"GoLandHomeworks/homeworks/Homework_22/configurations"
	"GoLandHomeworks/homeworks/Homework_22/redis_lib"
	"GoLandHomeworks/homeworks/Homework_22/start"
	"GoLandHomeworks/homeworks/Homework_22/users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	usersMongoStore, err := users.NewUsersStore(configurations.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "user_test22",
		CollectionName: "users",
	})
	if err != nil {
		log.Fatal(err)
	}
	redisStore, err := redis_lib.NewRedisConnect(configurations.RedisConfig{
		Host: "localhost",
		Port: "6379",
	})
	if err != nil {
		log.Fatal(err)
	}
	mainEndpoints := start.NewHttpEndpoints(usersMongoStore, redisStore)
	router.Methods("GET").Path("/main").HandlerFunc(mainEndpoints.Endpoint22())
	router.Methods("GET").Path("/main/{param}").HandlerFunc(mainEndpoints.Endpoint22WithParameter("param"))
	router.Methods("POST").Path("/main").HandlerFunc(mainEndpoints.PostEndpoint22())
	router.Methods("POST").Path("/register").HandlerFunc(mainEndpoints.RegisterEndpoint22())
	router.Methods("POST").Path("/login").HandlerFunc(mainEndpoints.LoginEndpoint22())
	router.Methods("GET").Path("/profile").HandlerFunc(mainEndpoints.ProfileEndpoint22())
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
