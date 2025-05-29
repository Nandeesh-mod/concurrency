package main

import (
	"fmt"
)

// to make the main go routine and boring go routing to communicate among themselves

func main() {
	c := make(chan string)
	go boaring("Boaring @", c)
	for item := range c {
		fmt.Println("You say : ", item)
	}

	fmt.Println("You are Boaring")
}

func boaring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%v : %d", msg, i)
		// time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
