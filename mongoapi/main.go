package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Roninanu777/mongo/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Sever is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on PORT 4000")
}
