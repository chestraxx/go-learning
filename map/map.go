package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Twitter"] = "https://twitter.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
