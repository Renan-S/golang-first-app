package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		emptyArray []int //This is technically a slice
		array1     [5]int
		array2     [5]int = [5]int{0, 1, 2, 3, 4}
	)

	array3 := [5]string{"i0", "i1", "i2", "i3", "i4"}
	lenByInitialValues := [...]int{10, 20, 30}

	fmt.Println(emptyArray, array1, array2, array3,
		lenByInitialValues, reflect.TypeOf(emptyArray))

	slice := []int{345, 234, 234, 896, 123} //It's a slice of an array
	fmt.Println(slice, reflect.TypeOf(slice))

	slice = append(slice, 12385)
	fmt.Println(slice)

	slice2 := array2[1:4] //Slice from index 1 to 4
	fmt.Println(slice2)

	array2[1] = 9 //Work as a pointer
	fmt.Println(slice2)

	//INTERNAL ARRAY
	fmt.Println("--------------------------")
	slice3 := make([]float64, 10, 11)             //Create new slice with type, length and size
	fmt.Println(slice3, len(slice3), cap(slice3)) //len for length and cap for size

	slice3 = append(slice3, 12, 13)
	fmt.Println(slice3, len(slice3), cap(slice3)) //Automatically doubles the size

	slice4 := make([]float32, 5) //You can pass only the length
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice4 = append(slice4, 6)
	fmt.Println(slice4, len(slice4), cap(slice4))
}
