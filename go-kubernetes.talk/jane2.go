package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", greetJohn)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func greetJohn(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:4000/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
