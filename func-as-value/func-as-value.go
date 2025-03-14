package main

import (
	"fmt"
)

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

	// closure
	numbers = []int{1, 2, 3}
	fmt.Println(numbers)

	factor := 5
	fifth := transform(&numbers, createTransformFunc(factor))
	fmt.Println(fifth)

	// recursion
	number := 5
	fmt.Println("number is", number, "| factorial is", factorial(number))

	// variadic func
	// numbers = []int{1, 2, 3}
	fmt.Println(1, 2, 3)
	sum_ := sum(1, 2, 3)

	fmt.Println("sum is", sum_)
}

func sum(numbers ...int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func createTransformFunc(factor int) transformFunc {
	return func(number int) int {
		return number * factor
	}
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
