package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in go");

	// var ptr *int;

	// fmt.Println("Value of pointer is ", ptr);

	myNumber := 23;

	var ptr1 = &myNumber;

	fmt.Println("value of pointer is ", ptr1);
	fmt.Println("value of pointer is ", *ptr1);

	*ptr1 /= 2;

	fmt.Println("value of pointer is ", *ptr1);
}