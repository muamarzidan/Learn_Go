package main

import "fmt"

// Suatu supermarket akan memberikan diskon dan cashback dengan beberapa kondisi berikut :
// 1. Jika pembeli berbelanja dengan harga lebih dari sama dengan Rp. 100000 , maka akan mendapatkan diskon sebanyak 10%

// 2. Jika pembeli berbelanja dengan harga lebih dari sama dengan Rp. 300000,  maka akan mendapatkan diskon sebanyak 20%

// 3. Jika pembeli adalah member dan berbelanja dengan harga lebih dari sama dengan Rp. 200000, maka akan medapatkan cashback sebanyak Rp. 25000

// 4. Jika pembeli adalah member dan berbelanja dengan harga lebih dari sama dengan Rp. 400000, maka akan mendapatkan cashback sebanyak Rp. 50000

// Deklarasikan tipe bentukan/ struct untuk menyimpan data pembeli dengan komponen boolean member dan harga belanja, kemudian buatlah program dalam bahasa GO untuk menghitung total biaya yang harus dibayar pembeli.

// Masukan berupa boolean member (true/false) , dan harga belanja yang bertipe bilangan bulat

// Keluaran berupa harga total yang harus dibayar yang bertipe bilangan real

// contoh input :
// true 450000
// contoh output :
// 310000

// contoh input :
// false 175000
// contoh output :
// 157500

// contoh input :
// true 230000
// contoh output :
// 182000

type pembeli struct {
	member bool
	harga  int
}

func main() {
	var p pembeli
	fmt.Scan(&p.member, &p.harga)
	hitungBiaya(p)
}

func hitungBiaya(p pembeli) {
	var total float64
	if p.member {
		if p.harga >= 400000 {
			total = float64(p.harga) - float64(p.harga)*0.2
			total = total - 50000
		} else if p.harga >= 20000 {
			total = float64(p.harga) - float64(p.harga)*0.1
			total = total - 25000
		}
	} else if p.member == false {
		if p.harga >= 300000 {
			total = float64(p.harga) - float64(p.harga)*0.2
		} else if p.harga >= 100000 {
			total = float64(p.harga) - float64(p.harga)*0.1
		}
	}
	fmt.Println(total)
}
