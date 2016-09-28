package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", sayHello)

	fmt.Println("Say hello to me! (on port :4000)")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request")
	fmt.Fprintf(w, "Hello, Jane!")
}
