package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	double := transform(&numbers, doubled)
	fmt.Println(double)

	triple := transform(&numbers, tripled)
	fmt.Println(triple)
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
