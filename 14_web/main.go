package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h>Hello World</h>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h>About US</h>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Starting server....")
	http.ListenAndServe(":3000", nil)
}
