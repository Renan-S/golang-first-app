package main

func main() {
	println(isApproved(7, 7))
}

func isApproved(number1, number2 float32) bool {
	defer func() {
		if recover := recover(); recover != nil { //Try to recover before PANIC
			println("Recover was a success!")
		}
	}()

	println("Starting average calculation")

	if average := (number1 + number2) / 2; average > 7 {
		return true
	} else if average < 7 {
		return false
	}

	panic("Average is 7") //This stops all the executions!
}
