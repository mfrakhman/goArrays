package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	website := map[string]string{
		"Google":             "http://www.google.com",
		"Amazon Web Service": "https://www.aws.com",
		"Vercel":             "https://www.vercel.com",
	}

	fmt.Println(website["Google"])

	website["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(website["LinkedIn"])
	delete(website, "Vercel")

	fmt.Println(website)

	coursePrice := make(floatMap, 3)

	coursePrice["Java"] = 100
	coursePrice["Python"] = 80
	coursePrice["Javascript"] = 90
	coursePrice["C#"] = 110

	fmt.Println(coursePrice)

	for k, v := range coursePrice {
		fmt.Println("key :", k)
		fmt.Println("value :", v)
	}
}
