package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is home route")
}

func routeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "this is method get")
	} else if r.Method == http.MethodPost {
		fmt.Fprintln(w, "this is method post")
	} else {
		fmt.Fprintln(w, "method not allowed")
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/handle", routeHandle)
	fmt.Println("listening on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
