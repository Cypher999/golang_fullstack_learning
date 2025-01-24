package main

import (
	"errors"
	"fmt"
)

func bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan 0")
	}
	return a / b, nil
}

func main() {
	hasil, err := bagi(9, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err)
	}
}
