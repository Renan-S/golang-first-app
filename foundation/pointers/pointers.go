package main

import "fmt"

func main() {
	//POINTERS ARE A MEMORY REFERENCE
	var (
		variable1        int  //Should be 0
		variable2        int  = 100
		emptyPointer     *int //Should be <nil>
		atributedPointer *int
	)

	// atributedPointer = variable1 - This operation is not possible because int != *int
	atributedPointer = &variable2 //Will print memory address

	fmt.Println(variable1, variable2, emptyPointer,
		atributedPointer, *atributedPointer)

	variable2 = 200

	fmt.Println(variable2, *atributedPointer)
}
