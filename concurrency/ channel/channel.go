package main

import "time"

func main() {
	channel := make(chan string) //chan for channel
	go write("Hello World!", channel)

	for message := range channel { //While channel is open the execution will wait
		println(message)
	}
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text //Send text into channel
		time.Sleep(time.Second)
	}

	close(channel)
}
