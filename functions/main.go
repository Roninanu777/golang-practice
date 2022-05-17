package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter("Roni")

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proRes, proStr := proAdder(2, 4, 5, 6, 7, 8)

	fmt.Println("Pro result is ", proRes, proStr)
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi"
}

func greeter(name string) {
	fmt.Printf("Namstey %v\n", name)
}
