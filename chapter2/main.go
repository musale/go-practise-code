package main

import "fmt"

func main() {
	fmt.Print("Enter temp in Fahreinheit: ")
	var fahreinheit float64
	fmt.Scanf("%f", &fahreinheit)

	celcius := (fahreinheit - 32) * 5 / 9

	fmt.Println("The temp in Celsius is ", celcius)
	metres := celcius * 0.3048
	fmt.Println("Assuming celsius is feet, then the metres are, ", metres)

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
	}

	nums := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := nums[0]
	for _, num := range nums {
		if smallest >= num {
			smallest = num
		}
	}

	fmt.Println("smallest number, ", smallest)
}
