package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// func as parameter value
	double := transform(&numbers, doubled)
	fmt.Println(double)

	triple := transform(&numbers, tripled)
	fmt.Println(triple)

	// func as return value
	moreNumbers := []int{6, 3, 4, 5}
	fmt.Println(moreNumbers)

	tranFunc := getTransformFunc(&moreNumbers)
	tNumber := transform(&moreNumbers, tranFunc)
	fmt.Println(tNumber)

	// anonymous func
	numbers = []int{1, 2, 3}
	fmt.Println(numbers)

	quad := transform(&numbers, func(number int) int {
		return number * 4
	})
	fmt.Println(quad)
}

func getTransformFunc(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return doubled
	} else {
		return tripled
	}
}

func transform(number *[]int, tranFunc func(int) int) []int {
	double := make([]int, len(*number))
	for i, v := range *number {
		double[i] = tranFunc(v)
	}

	return double
}

func doubled(number int) int {
	return number * 2
}

func tripled(number int) int {
	return number * 3
}
