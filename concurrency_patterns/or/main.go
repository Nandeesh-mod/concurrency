package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an interger")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers ", n)

	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting")
	end <- true
}

func gen(min int, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(time.Second * 5):
			fmt.Println("time.After() is called !!!")
		}
	}
}
