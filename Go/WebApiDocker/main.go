package main

import (
	"fmt"
	"net/http"
)

func rootPage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Root page. Check out /about")
}

func about(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "This is the about page.")
}

func requestHandler() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/about", about)

	err := http.ListenAndServe("localhost:8080", nil)
	fmt.Println(err)
}

func main() {
	requestHandler()
}