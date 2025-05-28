package main

import "fmt"

func main() {
	getStream := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	resStream := getStream()
	for item := range resStream {
		fmt.Println(item)
	}

}
