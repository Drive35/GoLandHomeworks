package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from main page, mate!")
}

func CalcPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	a := r.Form.Get("a")
	b := r.Form.Get("b")
	d := r.Form.Get("znak")
	aInt, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	total := 0
	if d == "+" {
		total = aInt + bInt
	} else if d == "-" {
		total = aInt - bInt
	} else if d == "*" {
		total = aInt * bInt
	} else if d == "/" {
		total = aInt / bInt
	}
	JsonResponse(w, total)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{
			Id:    1,
			Name:  "book",
			Price: 1500,
		},
		{
			2,
			"movie",
			2499,
		},
		{
			3,
			"pencil",
			100,
		},
	}
	JsonResponse(w, products)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	c := r.Form.Get("id")
	cInt, err := strconv.Atoi(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if cInt == 1 {
		JsonResponse(w, []Product{
			{
				Id:    1,
				Name:  "book",
				Price: 1500,
			},
		})
	} else if cInt == 2 {
		JsonResponse(w, []Product{
			{
				2,
				"movie",
				2499,
			},
		})
	} else if cInt == 3 {
		JsonResponse(w, []Product{
			{
				3,
				"pencil",
				100,
			},
		})
	} else if cInt > 3 {
		fmt.Fprint(w, "No data found")
	}
}

func JsonResponse(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(value)
}

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/calculate", CalcPage)
	http.HandleFunc("/products", GetProducts)
	log.Info("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
