package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {

	laptop := Product{
		Name:     "Mobile Phone",
		Price:    69.69,
		Quantity: 15,
	}

	fmt.Println("Product Name:", laptop.Name)
	fmt.Println("Price:", laptop.Price)
	fmt.Println("Quantity:", laptop.Quantity)
}
