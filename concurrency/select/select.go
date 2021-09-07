package main

import "time"

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		select { //Select makes channels autonomous. Channel1 doesn't have to wait for Channel2
		case messageCh1 := <-channel1:
			{
				println(messageCh1)
			}
		case messageCh2 := <-channel2:
			{
				println(messageCh2)
			}
		}
	}
}
