package main

import "fmt"

func main() {
	fmt.Println(calculations(10, 59))
}

func calculations(number1, number2 int) (sum, substraction int) {
	sum = number1 + number2
	substraction = number1 - number2
	return
}
