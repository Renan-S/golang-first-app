package main

import (
	"fmt"

	"github.com/paemuri/brdoc"
)

func main() {
	fmt.Println("Hello World! This is an example of how to use external package")
	fmt.Println(brdoc.IsCPF("063.532"))
}
