package product

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type HttpEndpoints interface {
	MainPage(w http.ResponseWriter, r *http.Request)
	TemplateExample(w http.ResponseWriter, r *http.Request)
	GetProductsList(w http.ResponseWriter, r *http.Request)
}
type productHttpEndpointsStruct struct {
	db ProductsStore
}

func NewProductHttpEndpoints(mongoConnect ProductsStore) HttpEndpoints {
	return &productHttpEndpointsStruct{db: mongoConnect}
}

func (ph *productHttpEndpointsStruct) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/main.html")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := r.Form.Get("name")
		price := r.Form.Get("price")
		description := r.Form.Get("description")
		if name == "" {
			http.Error(w, "please write name of product", http.StatusBadRequest)
			return
		}
		if price == "" {
			http.Error(w, "please write price of product", http.StatusBadRequest)
			return
		}
		if description == "" {
			http.Error(w, "please write description of product", http.StatusBadRequest)
			return
		}
		fmt.Println("form data ", name, price, description)
		http.ServeFile(w, r, "templates/main.html")
	}
}

func (ph *productHttpEndpointsStruct) TemplateExample(w http.ResponseWriter, r *http.Request) {
	templateFile, err := template.ParseFiles("templates/template.html")
	if err != nil {
		http.Error(w, "no html file", http.StatusInternalServerError)
		return
	}
	tp := &TemplateResponse{
		Title:       "Golang homework about template",
		Description: "template example",
		Price:       10000000,
	}
	if r.Method == "GET" {
		tp.Error = ""
		err = templateFile.Execute(w, tp)
		if err != nil {
			http.Error(w, "execute error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := r.Form.Get("name")
		price := r.Form.Get("price")
		p, _ := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal(err)
		}
		description := r.Form.Get("description")
		tp.Error = "Hello user"
		if name == "" {
			tp.Error = "please write name of product"
		}
		if price == "" {
			tp.Error = "please write price of product"
		}
		if description == "" {
			tp.Error = "please write description of product"
		}
		fmt.Println("data from product ", name, price, description)
		err = ph.db.AddProduct(Product{
			Name:        name,
			Price:       p,
			Description: description,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = templateFile.Execute(w, tp)
		if err != nil {
			http.Error(w, "execute error", http.StatusInternalServerError)
			return
		}
	}
}

func (ph *productHttpEndpointsStruct) GetProductsList(w http.ResponseWriter, r *http.Request) {
	templateFile, err := template.ParseFiles("templates/products_list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products, err := ph.db.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tp := &ListProductResponse{ProductField: products[0]}
	err = templateFile.Execute(w, tp)
	if err != nil {
		http.Error(w, "execute error", http.StatusInternalServerError)
		return
	}
}
