package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println(phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)

	doneChan <- true
}

func main() {
	done := make(chan bool)
	// go greet("Hello")
	// go greet("Hi")
	go slowGreet("Hello", done)
	// go greet("Hi")

	<-done
}
