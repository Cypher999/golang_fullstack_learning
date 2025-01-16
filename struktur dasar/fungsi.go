package main

import "fmt"

func salam(nama string) string {
	return "halo " + nama
}

func main() {
	fmt.Println(salam("Sandi"))
}
