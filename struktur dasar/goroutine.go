package main

import (
	"fmt"
	"time"
)

func tesLooping(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(1000)
	}
}

func main() {
	go tesLooping("pesan ini dari goroutine")
	tesLooping("pesan ini dari fungsi utama")
}
