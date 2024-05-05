package main

import "fmt"

func main() {
	var jariAlas, Tinggi float64
	fmt.Scanln(&jariAlas, &Tinggi)
	fmt.Println("Volume Tabung : ", Tabung(jariAlas, Tinggi))
	fmt.Println("Luas Tabung : ", luasTabung(jariAlas, Tinggi))
}

func Tabung(jariAlas, Tinggi float64) float64 {
	return 3.14 * jariAlas * jariAlas * Tinggi
}

func luasTabung(jariAlas, Tinggi float64) float64 {
	return 2 * 3.14 * jariAlas * (jariAlas + Tinggi)
}
