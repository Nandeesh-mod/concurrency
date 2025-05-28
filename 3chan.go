package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 20
		ch <- 30
		close(ch)
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// go func() {
	// 	fmt.Println(<-ch)
	// }()
}
