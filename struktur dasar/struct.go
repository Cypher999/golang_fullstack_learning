package main

import "fmt"

type Orang struct {
	nama    string
	umur    int
	menikah bool
}

func main() {
	sandi := Orang{
		nama:    "Sandi",
		umur:    26,
		menikah: true}
	fmt.Println("Halo " + sandi.nama)
	fmt.Println("Umur anda ", sandi.umur)
	if sandi.menikah {
		fmt.Println("anda sudah menikah")
	} else {
		fmt.Println("anda masih lajang")
	}
}
