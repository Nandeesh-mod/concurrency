package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for _, saluation := range[]string{"hello", "greetings", "good day"}{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println(saluation)
		}()
	}

	wg.Wait() 
}
