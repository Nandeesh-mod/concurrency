package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(time.Second * 5)
		close(c)
	}()
	fmt.Println("Blocking the read")

	select {
	case <-c:
		fmt.Println("Unblocked %v \n", time.Since(start))
	}
}
