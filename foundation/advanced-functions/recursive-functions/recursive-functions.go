package main

func main() {
	for index := uint(1); index <= 14; index++ {
		println(fibonacci(index))
	}
}

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
