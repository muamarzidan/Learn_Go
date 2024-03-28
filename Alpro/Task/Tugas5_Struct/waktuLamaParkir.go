package main

import "fmt"

// Definisikan tipe bentukan untuk menyatakan data waktu dengan komponen jam, menit, dan detik. Setelah itu buatlah program dalam bahasa Go untuk menghitung lama waktu kendaraan parkir.

// Masukan berupa 6 buah bilangan bulat yang menyatakan jam, menit, dan detik saat masuk parkir dan keluar parkir.

// Keluaran berupa 3 buah bilangan bulat yang menyatakan lama waktu parkir dalam jam, menit, dan detik. 


func main() {
	type parkir struct {
		jam int
		menit int
		detik int
	}

	var masuk, keluar parkir

	fmt.Scan(&masuk.jam, &masuk.menit, &masuk.detik)
	fmt.Scan(&keluar.jam, &keluar.menit, &keluar.detik)

	var lama parkir
	lama.detik = keluar.detik - masuk.detik
	lama.menit = keluar.menit - masuk.menit
	lama.jam = keluar.jam - masuk.jam

	if lama.detik < 0 {
		lama.detik += 60
		lama.menit--
	}

	if lama.menit < 0 {
		lama.menit += 60	
		lama.jam--
	}

	fmt.Println(lama.jam, lama.menit, lama.detik)
}