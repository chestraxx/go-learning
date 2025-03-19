package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)

	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Hello1", done)
	// dones[1] = make(chan bool)
	go greet("Hi1", done)
	// dones[2] = make(chan bool)
	go slowGreet("Hello 2 ", done)
	// dones[3] = make(chan bool)
	go greet("Hi 2 ", done)

	for range done {
		// fmt.Println(doneChan)
	}
}
