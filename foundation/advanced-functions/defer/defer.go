package main

func main() {
	defer function1() //function1 will be the last to run
	println(isApproved(7.2, 6.9))
	function2()
	function3()
	function4()
	function5()
}

func function1() {
	println("Execute func 1")
}

func function2() {
	println("Execute func 2")
}

func function3() {
	println("Execute func 3")
}

func function4() {
	println("Execute func 4")
}

func function5() {
	println("Execute func 5")
}

func isApproved(number1, number2 float32) bool {
	defer println("The average was calculated, returning value...") // This will be called always before return
	println("Starting average calculation")

	if average := (number1 + number2) / 2; average >= 7 {
		return true
	}
	return false
}
