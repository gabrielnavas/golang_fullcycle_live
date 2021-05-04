package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID   string
	Name string
}

type Products struct {
	Products []Product
}

func (p *Products) Add(product Product) {
	p.Products = append(p.Products, product)
}

func NewProduct() Product {
	return Product{
		ID: uuid.NewV4().String(),
	}
}
