// Buatlah sebuah program kalkulator sederhana dalam bahasa Go dengan menu pilihan penjumlahan, perkalian, dan pembagian.
// Lengkapi prosedur-prosedurnya dengan spesifikasi yang diminta.

// program kalkulator
// kamus
// 	pilih : integer
// algoritma
// 	repeat
// 		menu()
// 		output("Pilih (1/2/3/4)?")
// 		input (pilih)
// 		depend on (pilih)
// 			1: hitungJumlah()
// 			2: hitungkali()
// 			3: hitungBagi()
// 	until pilih == 4
// endprogram

// procedure menu()
// { IS:
// FS: Tercetak di layar tampilan sebagai berikut:
// MENU
// 1. Hitung Penjumlahan
// 2. Hitung Perkalian
// 3. Hitung Pembagian
// 4. Exit
// }

// procedure hitungJumlah()
// { IS:
// Proses: Membaca dua bilangan bulat dan menjumlahkan kedua bilangan itu
// FS: Mencetak hasil penjumlahan }

// procedure hitungkali()
// {IS: -
// Proses: Membaca dua bilangan bulat dan mengalikan kedua bilangan itu
// FS: Mencetak hasil perkalian }

// procedure hitungBagi()
// { IS: -
// Proses: Membaca dua bilangan real dan membagi bilangan pertama oleh bilangan kedua
// FS: Mencetak hasil pembagian }

// Contoh Masukan
// 1
// 11 12
// 2
// 34 45
// 3
// 1 2
// 4

// Contoh Keluaran
// ---------------------
// 		MENU
// ---------------------
// 1. Hitung Penjumlahan
// 2. Hitung Perkalian
// 3. Hitung Pembagian
// 4. Exit
// ---------------------
// Pilih (1/2/3/4)? 1
// Masukan dua bilangan yang akan dijumlahkan : 11 12
// Hasil penjumlahan : 23

// ---------------------
// 		MENU
// ---------------------
// 1. Hitung Penjumlahan
// 2. Hitung Perkalian
// 3. Hitung Pembagian
// 4. Exit
// ---------------------
// Pilih (1/2/3/4)? 2
// Masukan dua bilangan yang akan dikalikan : 34 45
// Hasil perkalian : 1530

// ---------------------
// 		MENU
// ---------------------
// 1. Hitung Penjumlahan
// 2. Hitung Perkalian
// 3. Hitung Pembagian
// 4. Exit
// ---------------------
// Pilih (1/2/3/4)? 3
// Masukan dua bilangan yang akan dibagi : 1 2
// Hasil pembagian : 0.5

// ---------------------
// 		MENU
// ---------------------
// 1. Hitung Penjumlahan
// 2. Hitung Perkalian
// 3. Hitung Pembagian
// 4. Exit
// ---------------------
// Pilih (1/2/3/4)? 4
// ---------------------

package main

import "fmt"

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)?")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		case 4:
			return
		}
	}
}

func menu() {
	fmt.Println("----------------------")
	fmt.Println("\tM E N U")
	fmt.Println("----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("----------------------")
}

func hitungJumlah() {
	var a, b int
	fmt.Print("Masukan dua bilangan yang akan dijumlahkan : ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil penjumlahan : ", a+b)
}

func hitungKali() {
	var a, b int
	fmt.Print("Masukan dua bilangan yang akan dikalikan : ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian : ", a*b)
}

func hitungBagi() {
	var a, b float64
	fmt.Print("Masukan dua bilangan yang akan dibagi : ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil pembagian : ", a/b)
}
