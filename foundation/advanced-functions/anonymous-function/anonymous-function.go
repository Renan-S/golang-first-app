package main

import "fmt"

func main() {

	func(text string) { //Automatic call
		fmt.Println(text)
	}("Hello world!")

	newVar := func(text string) string {
		return fmt.Sprintf("This is your text -> %s", text)
	}("Hello world!")

	fmt.Println(newVar)
}
