package main

import (
	"fmt"
)

type Product struct {
	id    string
	name  string
	price float64
}

func main() {
	// 1.
	hobbies := [3]string{"reading", "writing", "coding"}
	fmt.Println(hobbies)

	// 2.
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3.
	hob1 := hobbies[0:2]
	hob2 := hobbies[:2]
	fmt.Println(hob1, hob2)

	// 4.
	fmt.Println(hob1, len(hob1), cap(hob1))
	hob1 = hob1[1:]
	fmt.Println(hob1, len(hob1), cap(hob1))

	// 5.
	courseGoals := []string{"Learn Go", "Learn Python"}
	fmt.Println(courseGoals)

	// 6.
	courseGoals[1] = "Learn Java"
	courseGoals = append(courseGoals, "Learn C++")
	fmt.Println(courseGoals)

	// 7.
	products := []Product{
		{"1", "apple", 10.99},
		{"2", "banana", 9.99},
		{"3", "orange", 5.99},
	}
	products = append(products, Product{"4", "grape", 7.99})
	fmt.Println(products)
}
