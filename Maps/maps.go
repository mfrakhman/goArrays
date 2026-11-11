package main

import "fmt"

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
}
