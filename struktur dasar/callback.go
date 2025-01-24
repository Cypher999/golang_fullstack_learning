package main

import "fmt"

func anon_callback(x int, y int, f func(int, int)) {
	f(x, y)
}

func main() {
	anon_callback(2, 3, func(a int, b int) {
		fmt.Println(a * b)
	})
}
