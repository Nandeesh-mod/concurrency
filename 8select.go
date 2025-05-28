package main

import (
	"fmt"
	"time"
)

// No Channel is Ready to read data

func main() {
	var c <-chan interface{}
	select {
	// this case will never be unblocked because we are reading from nil channel
	case <-c:
		fmt.Println("Reading from c")
		// time.After returns a reciever channel from which we can read the time elapsed from duration
	case x := <-time.After(time.Second * 5):
		fmt.Println("Timed out", x)
	}
}
