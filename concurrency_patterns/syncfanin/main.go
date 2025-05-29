package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boaring(msg string) <-chan Message {
	ch := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			waitForIt := make(chan bool)
			ch <- Message{
				str:  fmt.Sprintf("%v : %v", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return ch
}

func fanin(ch1, ch2 <-chan Message) chan Message {
	ch := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			ch <- <-ch1
		}
	}()

	go func() {
		for i := 0; ; i++ {
			ch <- <-ch2
		}
	}()

	return ch
}

func main() {
	c := fanin(boaring("John"), boaring("Deo"))

	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1)

		msg2 := <-c
		fmt.Println(msg2)

		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You all are boaring I am leaving")
}
