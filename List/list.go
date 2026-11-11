package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	prices := [4]float64{10.2, 9.0, 8.45, 8.02}
	fmt.Println(prices)
	products := []Product{
		{id: "1", title: "One", price: 100},
		{id: "3", title: "Three", price: 200},
	}
	products = append(products, Product{id: "2", title: "Two", price: 100})
	fmt.Println(products)
}
