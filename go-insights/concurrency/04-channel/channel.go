package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre goroutines
// Channel é um tipo

// Producer function that sends data to a channel and then closes it
func producer(ch chan int) {
	defer close(ch)
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// Consumer function that receives data from a channel
func consumer(ch chan int) {
	for {
		val, ok := <-ch
		if !ok {
			// Channel is closed, exit the loop
			fmt.Println("Channel closed, exiting consumer")
			break
		}
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)

	fmt.Println("Done")
}
