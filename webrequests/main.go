package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", res)

	fmt.Println("Res is", res)

	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data is ", string(resp))
}
