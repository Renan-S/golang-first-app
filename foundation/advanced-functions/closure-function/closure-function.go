package main

func closure() func() {

	text := "Inside closure"

	function := func() {
		println(text)
	}

	return function
}

func main() {
	text := "Inside main"
	println(text)

	closureReturn := closure()
	closureReturn() //Maintain the scope from closure func
}
