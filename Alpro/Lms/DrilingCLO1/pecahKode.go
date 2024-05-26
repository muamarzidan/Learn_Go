package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)
	hitungJumlah(n, &jumlah)
	fmt.Println(jumlah)
}

func hitungJumlah(n int, jumlah *int) {
	/*  I.S terdefinisi bilangan bulat n yang menyatakan banyaknya bilangan
		F.S jumlah yang menyatakan hasil penjumlahan digit kedua dan digit keempat dari n bilangan */
		var i, bilangan int
		*jumlah = 0
		for i = 0; i < n; i++ {
			fmt.Scan(&bilangan)
			// digit1 := bilangan / 10000
			digit2 := (bilangan % 10000) / 1000
			// digit3 := (bilangan % 1000) / 100
			digit4 := (bilangan % 100) / 10
			// digit5 := bilangan % 10
			*jumlah += digit2 + digit4
		}
	}