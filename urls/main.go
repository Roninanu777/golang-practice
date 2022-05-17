package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sdf87sf78esd7f8"

func main() {
	fmt.Println("Welcome to urls in go")

	fmt.Println(myurl)

	//parsing url
	res, _ := url.Parse(myurl)

	fmt.Println(res.Scheme, res.Host, res.Path, res.Port(), res.RawQuery)

	qparams := res.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"], qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=roni",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
