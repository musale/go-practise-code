// What's the difference between a method and a function?
// a method is defined within a type (struct) and can be called using the dot (.)
// a function, well, is just a function. defined outside a struct scope but
// can be accessed globally(?). Man I'm bad at this :'D

// Why would you use an embedded anonymous field instead of a normal named field?
// embedded anonymous field allows us to access a struct and it's method
// within a different struct

package main

import (
	"fmt"
	"math"
)

func main() {
	// initialize circle and rectangle
	circle := Circle{0, 0, 4}
	rectangle := Rectangle{0, 0, 14, 13}

	fmt.Println("AREA Circle: ", circle.area())
	fmt.Println("AREA Rectangle: ", rectangle.area())

    fmt.Println("PERIMETER Circle: ", circle.perimeter())
    fmt.Println("PERIMETER Rectangle: ", rectangle.perimeter())
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (circle *Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (rectangle *Rectangle) perimeter() float64 {
	return rectangle.x1 + rectangle.y1 + rectangle.x2 + rectangle.y1
}

func (circle *Circle) perimeter() float64 {
	return math.Pi * circle.r * 2
}

type Shape interface {
	area() float64
	perimeter() float64 // calculates the perimeter of a Shape
}
