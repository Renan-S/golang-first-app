package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		println(<-write("Hello World!"))
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
