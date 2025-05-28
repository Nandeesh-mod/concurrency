package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Println(i)
		}(i)
	}
	fmt.Println("Unblocking go routings")
	close(begin)
	wg.Wait()
}
