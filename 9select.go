package main

import "fmt"

// what happens when no channel is ready, and
// we need to do something in the meantime?

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	select {
	case <-ch1:
		fmt.Println("Received on ch1")
	case <-ch2:
		fmt.Println("Received on ch2")
		// default runs when no channels is ready to receive data
	default:
		fmt.Println("Default")
	}
}
