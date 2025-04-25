package main

import "fmt"

type ValidationError struct {
	field string
	msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Kesalahan pada %s : %s", e.field, e.msg)
}

func cekNama(nama string) error {
	if len(nama) <= 0 {
		return ValidationError{
			field: "nama",
			msg:   "Nama Tidak Boleh Kosong",
		}
	}
	return nil
}

func main() {
	var nama string = "Sandi"
	err := cekNama(nama)
	if err == nil {
		fmt.Println("Halo " + nama)
	} else {
		fmt.Println(err.Error())
	}
}
