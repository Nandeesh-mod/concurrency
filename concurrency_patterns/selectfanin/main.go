package main

import "fmt"

func main() {

	right := make(chan int)
	end := right

	go func() {
		right <- 1
	}()

	for i := 0; i < 5; i++ {
		left := make(chan int)
		go func() {
			left <- 1 + <-right
		}()
		right = left
	}

	fmt.Println(<-end)

}
