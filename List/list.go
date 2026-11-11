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

	productName := make([]string, 2, 5)
	// make function creating new slice with 2 empty indexes

	productName[0] = "Apple"

	productName = append(productName, "Laptop")
	productName = append(productName, "Headphone")
	productName = append(productName, "Mouse")
	productName = append(productName, "Monitor")

	fmt.Println(productName)
	fmt.Println(productName[1])
	fmt.Println(productName[2])
	fmt.Println(productName[3])

	for i, val := range productName {
		fmt.Println("index:", i)
		fmt.Println("value:", val)
	}
}
