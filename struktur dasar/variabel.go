package main

import "fmt"

var nama string = "Sandi"
var tinggi_badang int = 165
var kecepatan_lari float64 = 2.5
var menikah bool = true
var angka [3]int = [3]int{1, 2, 3}
var mapping map[string]int = map[string]int{"a": 1, "b": 2}

const pi = 3.14

func main() {
	var lainnya string = "abc"
	umur := 26
	fmt.Println("Halo " + nama)
	fmt.Println(lainnya)
	fmt.Println("anda berumur ", umur, " tahun")
	fmt.Println("nilai pi ", pi)
	fmt.Println(angka)
	fmt.Println(mapping)
}
