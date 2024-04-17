// Buatlah program dalam bahasa Go untuk membaca nilai berupa bilangan bulat dan memasukkannya ke dalam array yang memiliki kapasitas maksimum 10,
// dengan mendeklarasikan sebuah tipe bentukan alias tabInt di bawah secara global:

// constant NMAX : integer = 10
// type tabInt : array[0..NMAX] of integer

// Setelah itu cetaklah semua elemen array itu sesuai template pencetakan. Pembacaan dan pencetakan bilangan dilakukan dengan prosedur bacaNilai dan cetakNilai.

// contoh masukan 
// 5
// 11 23 45 67 89
// contoh keluaran
// Nilai yang terdapat pada array: 11 23 45 67 89

// contoh masukan
// 3
// 1 2 3
// contoh keluaran
// Nilai yang terdapat pada array: 1 2 3

// contoh masukan
// 0
// contoh keluaran
// Array kosong 


package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var nilaiAkhir tabInt
	var banyakNilai int
	bacaNilai(&nilaiAkhir, &banyakNilai)
	cetakNilai(nilaiAkhir, banyakNilai)
}

func bacaNilai(NA *tabInt, n *int) {
	/*
		IS: Nilai akhir (NA) dan banyak elemen (n) terdefinisi sembarang,
		Proses: Membaca banyak elemen (n). Jika n > NMAX, maka n bernilai NMAX.
			Membaca nilai sebanyak n dan memasukkannya ke dalam array nilai akhir (NA).
		FS: Nilai akhir (NA) berisi nilai. Banyak elemen (n) berisi nilai NMAX, jika n > NMAX
	*/
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&NA[i])
	}
}

func cetakNilai(NA tabInt, n int) {
	/*
		IS: Nilai akhir (NA) dan banyak elemen (n) terdefinisi
		FS: Mencetak sebanyak n elemen array nilai akhir (NA)
			dengan format: "Nilai yang terdapat pada array: <n1> <n2> <n3> ..."
			Jika array kosong, maka cetak: "Array kosong"
	*/
	if n <= 0 {
		fmt.Println("Array kosong")
	} else {
		fmt.Print("Nilai yang terdapat pada array: ")
		for i := 0; i < n; i++ {
			fmt.Print(NA[i])
			if i < n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}