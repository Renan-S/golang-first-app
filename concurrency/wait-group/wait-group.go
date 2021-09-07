package main

import (
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		write("Hello World!")
		waitGroup.Done()
	}()

	go func() {
		write("The World doesn't care")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		println(text)
		time.Sleep(time.Second)
	}
}
