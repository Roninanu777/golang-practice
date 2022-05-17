package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go");

	var fruitList [4]string;
	fruitList[0] = "Apple";
	fruitList[1] = "Mango";
	fruitList[3] = "Peach";

	fmt.Println("Fruit list is ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"};
	fmt.Println("Vege list", len(vegList))
}