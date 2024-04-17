// Buatlah program dalam bahasa Go untuk membaca data berupa serangkaian bilangan dan memasukkannya ke dalam 
// array yang memiliki kapasitas maksimum 10 dengan mendeklarasikan sebuah tipe bentukan struktur tabInt di bawah secara global:

// constant NMAX : integer = 10
// type tabInt : array[0..NMAX] of integer

// Setelah itu cetaklah semua elemen array itu. Pembacaan dan pencetakan karakter dilakukan dengan prosedur baca dan cetak.

// Buatlah pula fungsi untuk menghitung jumlah dan rata-rata nilai yang terdapat pada array. Masukan dan keluaran pada program adalah sebagai berikut:

// Masukan berupa serangkaian bilangan bulat (positif dan negatif) yang diakhiri nilai 0 sebagai sentinel. 

// Keluaran berupa seluruh elemen array, jumlah, dan rata-ratanya. output berisi 2 baris, baris pertama berisi seluruh elemen array yang telah diurutkan,
// dan baris kedua berisi jumlah dan rata-rata dari elemen array tersebut. Jika jumlah elemen array melebihi kapasitas maksimum, maka hanya sebanyak NMAX elemen yang dicetak.

// contoh masukan 
// 10 2 23 -34 12 0
// contoh keluaran
// 10 2 23 12
// 81 16.2

// contoh masukan
// 0
// contoh keluaran 

// 0 NaN

// contoh masukan
// 1 2 3 -4 5 6 7 -8 9 10 -11 0
// contoh keluaran
// 1 2 3 5 6 7 9 10
// 55 5.5

package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dan n terdefinisi sembarang
		Proses: Setiap bilangan yang dibaca ditampung dalam sebuah variabel. 
			Jika bilangan negatif, maka ubah menjadi bilangan positif dan
			masukan ke dalam array. Pembacaan berakhir jika terbaca bilangan 0.
		FS: Array A sebanyak n elemen berisi bilangan positif
	*/
	var penampung int
	fmt.Scan(&penampung)

	for penampung != 0 {
		*n++
		if penampung < 0 {
			penampung = penampung * -1
		}

		if *n-1 < NMAX {
			A[*n-1] = penampung
		}

		fmt.Scan(&penampung)
	}

	if *n > NMAX {
		*n = NMAX
	}
}

func cetak(A tabInt, n int) {
	/*
		IS: Array A terdefinisi sebanya kn elemen
		FS: Array A tercetak di layar
	 */
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	/* Mengembalikan jumlah dari nilai bilangan pada elemen array A */
	var jumNilai int

	for i := 0; i < n; i++ {
		jumNilai = jumNilai + A[i]
	}

	return jumNilai
}

func rata_rata(A tabInt, n int) float64 {
	/* Mengembalikan rata-rata bilangan */
	var jumNilai int

	for i := 0; i < n; i++ {
		jumNilai = jumNilai + A[i]
	}

	return float64(jumNilai) / float64(n)
}