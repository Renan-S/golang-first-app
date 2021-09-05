package main

import "fmt"

func main() {
	add := add(13, 29)
	fmt.Println(add)

	var function = func(text string) {
		fmt.Println(text)
	}

	function("Please, print this")

	multipleReturns1, multipleReturns2 := mathFunctions(32, 58)
	fmt.Println(multipleReturns1, multipleReturns2)

	//OR

	_, multipleReturns3 := mathFunctions(32, 58)
	fmt.Println(multipleReturns3)
}

func add(number1 int8, number2 int8) int8 {
	return number1 + number2
}

func mathFunctions(number1, number2 int8) (int8, int16) { //Both params are int8
	return number1 + number2, int16(number1 * number2)
}
