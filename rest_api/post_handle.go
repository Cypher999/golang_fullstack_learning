package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is home route")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "method not allowed")
	} else {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "error when parsing data", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		fmt.Fprintln(w, "hello ", username)
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "error when parsing file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		if handler != nil {
			fmt.Fprintln(w, "filename is "+handler.Filename)
		}

	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/post", PostHandler)
	fmt.Println("listening on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
