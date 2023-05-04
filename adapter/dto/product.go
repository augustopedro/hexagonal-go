package dto

import "github.com/augustopedro/hexagonal-go/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (dto *Product) Bind(product *application.Product) (*application.Product, error) {
	if dto.ID != "" {
		product.ID = dto.ID
	}
	product.Name = dto.Name
	product.Price = dto.Price
	product.Status = dto.Status
	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}
