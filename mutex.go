package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	var oncea sync.Once
	var onceb sync.Once

	var increments sync.WaitGroup

	increment := func() {
		count++

	}

	// decrement := func() {
	// 	count--

	// }

	increments.Add(2)
	go func() {
		oncea.Do(increment)
		increments.Done()
	}()

	go func() {
		onceb.Do(increment)
		increments.Done()
	}()

	increments.Wait()
	fmt.Println(count)
}
