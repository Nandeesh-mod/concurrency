package main

import "fmt"

// channels are the first class values just like strings or integers

func main() {
	// If I want boaring function to call twice

	// here teddy and jedi are independent
	td := Boaring("teddy")
	jd := Boaring("Jedi")
	c := fanin(td, jd)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("I am Leaving")

}

func fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- <-input1
		}
	}()

	go func() {
		for i := 0; ; i++ {
			c <- <-input2
		}
	}()
	return c
}

func Boaring(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%v : %v", msg, i)
		}
	}()

	return ch
}
