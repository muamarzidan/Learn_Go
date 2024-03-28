package main

import "fmt"

// Deklarasikan tipe bentukan untuk menampung data pemain bola dengan komponen dan contoh data sebagai berikut:

// NAMA GOL ASSISTS
// A    20   13
// B    14   4
// C    32   3

// Kemudian buatlah program dalam bahasa Go untuk menghitung rata-rata gol dan assist dari ketiga pemain.

// Masukan terdiri dari 3 baris data pemain dengan masing-masing barisnya berisi sebuah string (nama) dan dua buah bilangan bulat (gol dan assist).

// Keluaran berupa rata-rata gol dan assist dari ketiga pemain.


func main() {
	type pemain struct {
		nama   string
		gol    int
		assist int
	}

	var a, b, c pemain

	fmt.Scan(&a.nama, &a.gol, &a.assist)
	fmt.Scan(&b.nama, &b.gol, &b.assist)
	fmt.Scan(&c.nama, &c.gol, &c.assist)

	var rataGol, rataAssist float64
	rataGol = float64(a.gol+b.gol+c.gol) / 3
	rataAssist = float64(a.assist+b.assist+c.assist) / 3

	fmt.Println(rataGol)
	fmt.Println(rataAssist)
}
