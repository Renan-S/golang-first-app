package main

import (
	"fmt"
	"time"
)

func main() {
	index := 0
	for index < 10 { //Similar as a while
		index++
		fmt.Println("Index going plus 1", index)
		time.Sleep(time.Second) //Sleep for 1 second
	}

	for jules := 0; jules <= 10; jules++ {
		fmt.Println("Look at him go", jules)
	}

	names := []string{"John", "Doe", "Renan"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range names { //If you don't wanna use the index
		fmt.Println(name)
	}

	for index, letter := range "WORDS" {
		// fmt.Println(index, letter) Will print ASCII number
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name": "Renan",
		"age":  "31",
	}

	for _, element := range user {
		fmt.Println(element)
	}

	// for { For infinite loop just declare for

	// }
}
