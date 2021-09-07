package main

import "time"

func main() {
	go write("Hello World!") //Don't wait for the first write to finish
	write("The World doesn't care")
}

func write(text string) {
	for { //Infinite loop
		println(text)
		time.Sleep(time.Second)
	}
}
