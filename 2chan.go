package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// create a buffer which stores byte
	var stdBuffer bytes.Buffer

	// during the exit of function write the data in buffer to the standerd out
	defer stdBuffer.WriteTo(os.Stdout)

	// created a bufferd channel with capacity 4
	inStream := make(chan int, 2)

	go func() {
		// close the channel once done writing
		defer close(inStream)

		// write the result `Producer Done` to the buffer
		defer fmt.Fprintln(&stdBuffer, "Producer Done")

		//iterate over 0 to 5 and write the value to instream channel concurrently
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdBuffer, "Sending %d\n", i)
			inStream <- i
		}
	}()

	// writing the received data to buffer
	for integer := range inStream {
		fmt.Fprintf(&stdBuffer, "Received : %d", integer)
	}
}

// programm to pass the data to stream and receive from stream and write the log to the
// buffer and finally write the bufferd data to the output
