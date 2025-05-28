package main

import "fmt"

func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}

	select {
	// if c1 send any data than the code under this is executed
	case <-c1:
		fmt.Println("data sent to c1")
		// Do Something
	case <-c2:
		fmt.Println("data sent to c2")
		// Do Something
	case c3 <- struct{}{}:
		fmt.Println("data received to c3")
		// Do somethings
	}
}
