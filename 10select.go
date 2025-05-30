package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})

	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()

	workCount := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCount++
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCount)
}
