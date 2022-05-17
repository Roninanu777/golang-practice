package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go")

	//var fruitList = []string{"Apple", "Mango", "Peach", "Banana", "Plum"};

	// fruitList = append(fruitList[:2]);

	highScores := make([]int, 4);
	
	highScores[0] = 234;
	highScores[1] = 577;
	highScores[2] = 124;
	highScores[3] = 987;
	// highScores[4] = 555;

	highScores = append(highScores, 222,333,444)


	sort.Ints(highScores);

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "python", "ruby", "swift"};
	fmt.Println(courses);

	var index int = 2;
	fmt.Println(append(courses[:index], courses[index+1:]...))
}