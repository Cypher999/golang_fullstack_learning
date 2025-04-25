package main

import "fmt"

// Fungsi untuk menangkap panic menggunakan recover
func tangkapPanic() {

	// Menyebabkan panic
	panic("Terjadi kesalahan fatal!")
}

func main() {
	fmt.Println("Program dimulai")

	// Memanggil fungsi dengan panic
	tangkapPanic()

	fmt.Println("Program tetap berjalan setelah recover")
}
