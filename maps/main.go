package main

import "fmt"

func main() {
	fmt.Println("welcome to maps in go");

	languages := make(map[string]string);

	languages["JS"] = "Javascript";
	languages["RB"] = "Ruby";
	languages["PY"] = "Python";

	
	fmt.Println("JS shorts for: ", languages["JS"]);

	delete(languages, "RB");
	fmt.Println("List of all lang", languages);

	// loops in map
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value);
	}
}