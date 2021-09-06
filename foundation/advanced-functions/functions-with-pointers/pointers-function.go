package main

func main() {
	number := 10
	println(number)
	invertSignal(&number)
	println(number)
}

func invertSignal(number *int) {
	//This should be used with care because of the colateral damage
	*number = *number * -1 //This will change values directly where the variable is stored
}
