package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	college string
	major   string
}

func main() {
	person1 := person{"Jonh", "Doe", 31, 170}
	fmt.Println(person1)

	student1 := student{person1, "MIT", "Computer Science"}
	fmt.Println(student1.name, student1.major)
}
