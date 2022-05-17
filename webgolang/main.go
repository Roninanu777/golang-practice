package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verd video")
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	var myurl string = "http://localhost:8000/get"

	res, err := http.Get(myurl)
	checkNilErr(err)

	defer res.Body.Close()

	fmt.Println("Status code: ", res.StatusCode, res.Status)
	fmt.Println("Content length is: ", res.ContentLength)

	var resStr strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	byteCount, _ := resStr.Write(content)
	fmt.Println(byteCount)
	fmt.Println(resStr.String())

	//fmt.Println(string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learncodeonline.in"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	checkNilErr(err)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)

	checkNilErr(err)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}

	data.Add("firstname", "Roni")
	data.Add("middlename", "Raj Kamal")
	data.Add("lastname", "Pradhan")
	data.Add("email", "roni@roni.dev")

	res, err := http.PostForm(myurl, data)
	checkNilErr(err)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)

	checkNilErr(err)

	fmt.Println(string(content))
}
