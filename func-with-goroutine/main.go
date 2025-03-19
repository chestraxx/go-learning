package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println(phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
}

func main() {
	greet("Hello")
	greet("Hi")
	slowGreet("Hello")
	greet("Hi")
}
