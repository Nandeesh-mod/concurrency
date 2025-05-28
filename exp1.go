package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	var data int
	var memoryAccess sync.Mutex

	go func(){
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	time.Sleep(time.Second * 1)	
	memoryAccess.Lock()
	if data == 0{
		fmt.Println("value of data is : ",  data)
	}else{
		fmt.Println("value of data is : ", data)
	}
	memoryAccess.Unlock()
}
