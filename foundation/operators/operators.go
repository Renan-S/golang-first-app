package main

import "fmt"

func main() {
	//ARITHMETICS
	add := 1 + 2
	subtract := 3 - 4
	division := 12 / 2
	multiplication := 10 * 5
	remainder := 10 % 2

	fmt.Println(add, subtract, division, multiplication, remainder)

	//RELATIONAL
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//LOGICAL
	fmt.Println("----------------------------")
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(!false)

	//UNARY
	number := 10
	number++
	number += 14
	fmt.Println(number)

	number--
	number -= 14
	fmt.Println(number)

	number *= 10
	fmt.Println(number)
}
