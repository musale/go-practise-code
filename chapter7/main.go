package main

import "fmt"

func main() {
	// 1
	xs := []int{1, 2, 3}
	fmt.Println("Sum: ", sum(xs...))

	// 2
	fmt.Println("ODD ", halve(1))
	fmt.Println("EVEN ", halve(2))
}

func sum(args ...int) int {
	// sum is a function which takes a slice of numbers and adds them together.
	// What would its function signature look like in Go?
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func halve(args int) (int, bool) {
	// Write a function which takes an integer and halves it and returns true if
	// it was even or false if it was odd. }

	half = args / 2

	if half % 2 == 0 {
		// even
		return half, true
	} else {
		// odd
		return half, false
	}
}
