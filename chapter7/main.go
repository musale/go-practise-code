package main

import "fmt"

func main() {
	// 1
	xs := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println("Sum: ", sum(xs...))

	// 2
	fmt.Println(halve(1))
	fmt.Println(halve(2))

	// 3
	fmt.Println(greatestAlive(xs...))

	// 4
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	// 5
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(7))

	// 6.
	fmt.Println(defs())
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

	half := args / 2

	if half%2 == 0 {
		// even
		return half, false
	} else {
		// odd
		return half, true
	}
}

func greatestAlive(args ...int) int {
	// Write a function with one variadic parameter that finds the greatest
	// number in a list of numbers.

	greatest := args[0]
	for _, value := range args {
		if value >= greatest {
			greatest = value
		}
	}

	return greatest
}

func makeOddGenerator() func() uint {
	//  write a makeOddGenerator function that generates odd numbers.
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func fib(args int) int {
	//  Write a recursive function which can find fib(n)
	if args == 0 || args == 1 {
		return args
	}
	return fib(args-1) + fib(args-2)
}

func defs() string {
	definitions := "\n a defer is a schedule for a function to be called when the function finishes running.\n a panic is a programmer error that is not easy to recover from.\n a recover is a way used to handle panic errors."
	return definitions
}
