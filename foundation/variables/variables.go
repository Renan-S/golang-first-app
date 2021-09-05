package main

import "fmt"

func main() {
	var variable1 string = "variable1" //Declared type
	fmt.Println(variable1)

	variable2 := "variable2" //Type inference
	fmt.Println(variable2)

	var ( //Multiple declared types var
		variable3 string = "variable3"
		variable4 string = "variable4"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "variable5", "variable6" //Multiple type inference
	fmt.Println(variable5, variable6)

	const const1 string = "const1" //Constant
	fmt.Println(const1)

	const ( //Multiple constants
		const2 string = "const2"
		const3 string = "const3"
	)
	fmt.Println(const2, const3)

	variable5, variable6 = variable6, variable5 //Invert values
	fmt.Println(variable5, variable6)
}
