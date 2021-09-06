package main

import "fmt"

func main() {
	fmt.Println(sum(1, 23, 59, 283, 2828349, 92))

}

func sum(firstNum, secondNum int, numbers ...int) (int, int, int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return firstNum, secondNum, total
}
