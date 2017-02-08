// How do you get the memory address of a variable?
// using the & sign. i.e. &x

// How do you assign a value to a pointer?
// by using a * before the pointer name i.e. *x = 7

// How do you create a new pointer?
// by using * before type of stored value i.e. *int

package main

import "fmt"

func square(x *float64) {
	*x = *x * *x

    // 1
    fmt.Println(*x) // 2.25
}
func main() {
	m := 1.5
    y := 2
    x := 1
	square(&m)
    swap(&x, &y)
}

func swap(xPtr *int, yPtr *int)  {
    // Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)
    // should give you x=2 and y=1).

    tmp := *xPtr
    *xPtr = *yPtr
    *yPtr = tmp
    fmt.Println("x1 ", *xPtr)
    fmt.Println("y2 ", *yPtr)
}
