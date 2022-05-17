package main

import "fmt"

func main() {
	fmt.Printf("Welcome to structs in golang")

	roni := User{"Roni", "roni@go.dev", true, 29}
	fmt.Println(roni)
	fmt.Printf("Roni details are: %+v\n", roni)
	fmt.Printf("Name is %v and email is %v.\n", roni.Name, roni.Email)

	roni.GetStatus()
	roni.NewMail()

	fmt.Printf("Name is %v and email is %v.\n", roni.Name, roni.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "roni@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
