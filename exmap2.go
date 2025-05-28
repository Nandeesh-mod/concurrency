package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	fmt.Println("Calling go progamm")
	wg.Add(1)
	go hello(&wg)
	wg.Wait()

}


func hello(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("HGello")
}
