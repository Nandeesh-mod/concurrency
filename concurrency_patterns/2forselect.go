package main

import "fmt"

func main() {
	done := make(chan interface{})
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for _, i := range []string{"a", "b", "c", "d", "e"} {

			select {
			case <-done:
				return
			case stringStream <- i:
			}
		}
	}()
	for i := range stringStream {
		fmt.Println(i)
	}

}
