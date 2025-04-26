package main

import (
	"fmt"
	"net/http"
)

func helloWorldHAndler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", helloWorldHAndler)
	fmt.Println("Server running at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
