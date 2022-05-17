package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")

	content := "This need to go in a file - roni.dev"

	file, err := os.Create("./mydevfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mydevfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
