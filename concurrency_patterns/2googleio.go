package main

import "fmt"

// channels are the first class values just like strings or integers

func main() {
	// If I want boaring function to call twice

	td := Boaring("Teddy")
	jd := Boaring("Jeddy")

	for i := 0; i < 5; i++ {
		fmt.Println("You say : ", <-td)
		fmt.Println("You say : ", <-jd)

		// If jeddy is more talkitve than teddy then this wont work , we can use fan in pattern

		// here even though jeddy received value it is blocked by <-td until td received value
	}

	fmt.Println("I am Leaving")

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
