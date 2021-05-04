package main

import (
	"fmt"

	"github.com/gabrielnavas/go_mod_example/http"
	"github.com/gabrielnavas/go_mod_example/model"
)

func main() {
	product1 := model.NewProduct()
	product1.Name = "Ferray"

	product2 := model.NewProduct()
	product2.Name = "House"

	products := model.Products{}
	products.Add(product1)
	products.Add(product2)

	server := http.NewWebServer()
	server.Products = products.Products
	server.Serve()

	fmt.Println(products)
}
