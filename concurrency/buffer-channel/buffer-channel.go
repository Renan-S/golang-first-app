package main

func main() {
	channel := make(chan string, 2) //Buffer defines the capacity of the channel
	channel <- "Hello World!"       //This is trying to send a value for without someone waiting to get
	channel <- "World doesn't care!"

	println(<-channel, <-channel)
}
