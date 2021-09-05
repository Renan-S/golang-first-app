package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	var emptyUser user
	fmt.Println(emptyUser)

	newUser := user{"Renan", 31, address{"Nothing hill", 123}}
	fmt.Println(newUser)

	pickValues := user{age: 31}
	fmt.Println(pickValues)
}
