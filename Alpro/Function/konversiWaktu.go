package main

import "fmt"

func main() {
	var jam, menit, detik int

	fmt.Scanln(&jam, &menit, &detik)
	fmt.Println("Konversi Waktu =", konversiWaktu(jam, menit, detik))
}

func konversiWaktu(jam, menit, detik int) int {
	return (jam * 3600) + (menit * 60) + detik
}
