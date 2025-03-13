package main

import "fmt"

type floatMap map[string]float64

func main() {
	// map
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

	// make func
	userNames := make([]string, 2, 5)
	userNames[0] = "John"
	userNames[1] = "Jane"
	fmt.Println(userNames)

	// make func with map
	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.9
	courseRatings["Python"] = 4.8
	courseRatings["JavaScript"] = 4.7
	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
