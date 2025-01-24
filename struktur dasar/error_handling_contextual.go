package main

import "fmt"

func cekUmur(a int) error {
	if a < 18 {
		return fmt.Errorf("Umur anda belum memenuhi syarat")
	}
	return nil
}
func main() {
	err := cekUmur(12)
	if err == nil {
		fmt.Println("umur anda memenuhi syarat")
	} else {
		fmt.Println(err)
	}

}
