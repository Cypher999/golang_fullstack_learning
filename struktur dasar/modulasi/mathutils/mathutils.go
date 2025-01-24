package mathutils

import "math"

func Akar(angka float64, eksponen float64) float64 {
	return math.Pow(angka, (1 / eksponen))
}

func LuasSegitiga(alas float64, tinggi float64) float64 {
	return (alas * tinggi) / 2
}

func LuasLingkaran(jari float64) float64 {
	return 3.14 * (math.Pow(jari, 2))
}
