package main

import (
	"fmt"
	"time"
)

// programm which pass integer from 1 to 10 to channel and close it once it send all values

func main() {
	ch := make(chan int, 2)
	go func() {
		ch <- 0
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Printf("reading 4")
		ch <- 4
		fmt.Printf("Blocking")
		ch <- 5
		fmt.Printf("Blocking")

		close(ch)
	}()

	for item := range ch {
		fmt.Println(item)
	}

	time.Sleep(1 * time.Second)
}
