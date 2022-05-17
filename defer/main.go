package main

import "fmt"

func main() {
	myDefer()
	defer fmt.Println(" in go")
	defer fmt.Println("defer")

	fmt.Println("Welcome to ")

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
