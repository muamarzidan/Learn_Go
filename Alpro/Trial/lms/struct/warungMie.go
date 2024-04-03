// Jika mie tersebut pedas, maka harga bertambah Rp. 1000.
// Topping yang sama dapat ditambahkan lebih dari sekali.

// Terdapat suatu warung mie yang menjual mie dengan berbagai topping. Definisikan tipe data untuk menyimpan informasi tentang
// semangkok mie dengan komponen string tipe mie, boolean untuk menandakan apakah pedas atau tidak, dan string topping. Selanjutnya,
// buatlah program dalam bahasa Go untuk menghitung total harga yang harus dibayar.

// Masukan berupa string tipe mie, boolean pedas atau tidak, dan input topping terus menerus hingga menginputkan "0" atau topping lebih dari 5.

// Keluaran berupa total harga yang bertipe bilangan bulat.

// Catatan:
// Terdapat 3 tipe mie dengan harga yang berbeda:
// harga "kwetiau"  Rp. 8000,
// harga "bihun" Rp. 7000,
// harga "kuning" Rp. 9000.

// Terdapat 4 topping dengan harga yang berbeda:
// harga "ayam" Rp. 5000,
// harga "telur" Rp. 3000,
// harga "sayur" Rp. 3000 ,
// harga "pangsit" Rp. 5000.

// Jika mie tersebut pedas, maka harga bertambah Rp. 1000.
// Topping yang sama dapat ditambahkan lebih dari sekali.

// contoh masukan
// kwetiau true
// 0
// contoh keluaran
// 9000

// contoh masukan
// bihun false == 7000
// ayam == 5000
// ayam == 5000
// ayam == 5000
// sayur == 3000
// telur == 3000
// contoh keluaran
// 26000

// contoh masukan
// kuning true == 10000
// ayam == 5000
// sayur == 3000
// telur == 3000
// 0
// contoh keluaran
// 21000

// package main

// import "fmt"

// type mie struct {
// 	tipe    string
// 	pedas   bool
// 	topping string
// }

// func main() {
// 	var m mie
// 	var iterasi, total int

// 	fmt.Scan(&m.tipe, &m.pedas)
// 	fmt.Scan(&m.topping)

// 	if m.topping == "0" {
// 		hitungHarga(m, &total)
// 		fmt.Println(total)
// 		return
// 	} else {
// 		iterasi++
// 		for iterasi < 5 {
// 			if m.topping == "0" {
// 				break
// 			}
// 			hitungHarga(m, &total)
// 			fmt.Scan(&m.topping)
// 			iterasi++
// 		}
// 		fmt.Println(total)
// 	}
// }

// func hitungHarga(m mie, totalBayar *int) {
// 	var harga int

// 	if m.tipe == "kwetiau" {
// 		harga = 8000
// 	} else if m.tipe == "bihun" {
// 		harga = 7000
// 	} else if m.tipe == "kuning" {
// 		harga = 9000
// 	}

// 	if m.pedas {
// 		harga += 1000
// 	}

// 	if m.topping == "ayam" {
// 		harga += 5000
// 	} else if m.topping == "telur" {
// 		harga += 3000
// 	} else if m.topping == "sayur" {
// 		harga += 3000
// 	} else if m.topping == "pangsit" {
// 		harga += 5000
// 	}

// 	*totalBayar += harga
// }

// 	for {
// 		fmt.Scan(&m.tipe, &m.pedas)
// 		if m.tipe == "0" {
// 			break
// 		}
// 		fmt.Scan(&m.topping)

// 		total += hitungHarga(m)
// 	}
// 	fmt.Println(total)
// }

package main

import "fmt"

type mie struct {
	tipe    string
	pedas   bool
	topping string
}

func main() {
	var m mie
	var total int

	fmt.Scan(&m.tipe, &m.pedas)

	hitungHarga(m, &total)
	fmt.Println(total)
}

func hitungHarga(m mie, totalBayar *int) {
	var harga, iterasi int

	if m.tipe == "kwetiau" {
		harga = 8000
	} else if m.tipe == "bihun" {
		harga = 7000
	} else if m.tipe == "kuning" {
		harga = 9000
	}

	if m.pedas {
		harga += 1000
	}

	for iterasi = 0; iterasi < 5; iterasi++ {
		fmt.Scan(&m.topping)
		if m.topping == "0" {
			break
		} else if m.topping == "ayam" {
			harga += 5000
		} else if m.topping == "telur" {
			harga += 3000
		} else if m.topping == "sayur" {
			harga += 2000
		} else if m.topping == "pangsit" {
			harga += 5000
		}
	}

	*totalBayar += harga
}
