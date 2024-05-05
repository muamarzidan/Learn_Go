// sebuah fungsi digunakan untuk menghitung total promo yang diberikan oleh swalayan ABC. 
// Promo diskon dan cashback diberikan apabila total belanja pembeli habis dibagi 100 atau habis dibagi n tetapi tidak habis dibagi 100, dengan n untuk diskon adalah 4 sedangkan untuk cashback adalah 3.

// Fungsi dipanggil dalam program dengan masukan dan keluaran sebagai berikut:

// Masukan berupa sebuah bilangan bulat yang menyatakan total belanja pembeli.
// Keluaran adalah 0 bila tidak memenuhi kondisi atau 1 bila memenuhi kondisi.


package main

import "fmt"

func main() {
	var banyaknyaPembeli, totalBelanja, i int
	var DapatDiskon, DapatCashback int
	fmt.Scan(&banyaknyaPembeli)
	DapatDiskon = 0
	DapatCashback = 0
	for i = 1; i <= banyaknyaPembeli; i++ {
		fmt.Scan(&totalBelanja)
		DapatDiskon = DapatDiskon + diskon(totalBelanja)
		DapatCashback = DapatCashback + cashback(totalBelanja)
	}
	fmt.Println(DapatDiskon,DapatCashback)
}

func diskon(totbel int) int {

	/*Mengembalikan angka 1 jika belanja 
	pembeli memenuhi syarat diskon dan 
	angka 0 jika tidak memenuhi syarat*/
		if totbel % 100 == 0 || (totbel % 4 == 0 && totbel % 100 != 0) {
			return 1
		} else {
			return 0
		}
	}
	
	
	func cashback(totbel int) int {
	
	/*Mengembalikan angka 1 jika belanja pembeli 
	memenuhi syarat cashback dan 
	angka 0 jika tidak memenuhi syarat*/
		if totbel % 100 == 0 || (totbel % 3 == 0 && totbel % 100 != 0) {
			return 1
		} else {
			return 0
		}
	}