package main

import "fmt"

// The first question of multiple channels being ready simultaneously seems interesting.
// Let’s just try it and see what happens!

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i > 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("Total c1 :%d and c2 : %d", c1Count, c2Count)
}
