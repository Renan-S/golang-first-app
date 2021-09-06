package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

func (rec rectangle) area() float64 {
	return rec.height * rec.width
}

func (circ circle) area() float64 {
	return math.Pi * math.Pow(circ.radius, 2) //Pi * radius squared
}

func writeArea(localForm form) {
	fmt.Printf("The area is %0.2f\n", localForm.area())
}

func main() {
	rectangle := rectangle{height: 12, width: 25}
	writeArea(rectangle)

	circle := circle{14}
	writeArea(circle)
}
