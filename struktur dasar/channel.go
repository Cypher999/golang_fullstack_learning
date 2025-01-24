package main

import "fmt"

func worker(ch chan string) {
	ch <- "tugas selesai"
}
func main() {
	ch := make(chan string)
	go worker(ch)
	msg := <-ch
	fmt.Println(msg)
}
