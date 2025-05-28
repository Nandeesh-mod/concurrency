package main

import (
	"fmt"
	"sync"
)

func main() {

	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new instance")
			return 10
		},
	}
	ins := myPool.Get() // once
	fmt.Println(ins)
	myPool.Put(ins)
	myPool.Put(ins)

	_ = myPool.Get()
	_ = myPool.Get()
	_ = myPool.Get()

}
