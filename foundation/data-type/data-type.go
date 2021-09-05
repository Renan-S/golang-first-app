package main

import (
	"errors"
	"fmt"
)

func main() {
	//INTEGER
	var ( //Integer data types
		dynamicInt    int   = 1231231231231312312 //Use SO basic architecture
		lowInt        int8  = -123
		byteInt       byte  = 123 //int8 alias/ can't be negative
		mediumInt     int16 = -12312
		highMediumInt int32 = 1231231231
		runeInt       rune  = -1231231231 //int32 alias
		bigInt        int64 = 1231231231231231231
		unassignedInt uint8 = 123 //Can't be negative
	)

	inferenceInt := 1234
	fmt.Println(dynamicInt, lowInt, byteInt, mediumInt, highMediumInt,
		runeInt, bigInt, unassignedInt, inferenceInt)

	//FLOAT
	var ( //Float data types
		lowFloat  float32 = -123.3284
		highFloat float64 = 123134939493.345
	)

	inferenceFloat := 12345.53
	fmt.Println(lowFloat, highFloat, inferenceFloat)

	//STRING
	var string1 string = "this is a string" //There is no CHAR on Go
	string2 := "this a second string"
	char := 'B' //ASCII table number
	fmt.Println(string1, string2, char)

	//INITIAL VALUES
	var emptyString string
	var emptyNum int     //0
	var emptyBool bool   //false
	var emptyError error //<nil> can be know as null too
	fmt.Println(emptyString, emptyNum, emptyBool, emptyError)

	//BOOLEAN
	var boolean bool = true
	boolean2 := false
	fmt.Println(boolean, boolean2)

	//ERROR
	var newError error = errors.New("this is a new error")
	fmt.Println(newError)
}
