package main

import "fmt"

func main() {
	var kata string
	var jumlah, index int
	index = 0
	fmt.Scanln(&kata)
	jumlah = hitungLowercase(kata, index)
	fmt.Print(jumlah)
}

func hitungLowercase(kata string, index int) int {
	/*{ I.S. Terdefinisi parameter kata sebagai input string
	F.S. menampilkan jummlah  jumlah huruf lowercase dari kata menggunakan fungsi rekursif*/
	if index == len(kata) { // masukan kondisi untuk mengecek index dengan jumlah string
		return 0
	}
	if 'a' <= kata[index] && kata[index] <= 'z' {
		return 1 + hitungLowercase(kata, index+1) // gunakan fungsi rekursif yang tepat
	}
	return hitungLowercase(kata, index+1)
}