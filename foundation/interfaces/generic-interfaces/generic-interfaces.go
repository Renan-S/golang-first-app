package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic(123)
	generic(true)
	generic("This is a string")

	genericMap := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     int64(2353435),
	}

	fmt.Println(genericMap)
}
