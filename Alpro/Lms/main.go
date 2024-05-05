// Buatlah sebuah prosedur rekursif yang mencetak bentuk biner dari bilangan bulat positif N. 

// Masukan berupa sebuah bilangan bulat positif N. 
// Keluaran mencetak bentuk biner dari N. Jadi, bila N = 11, maka akan tercetak 1011.

// contoh
// input
// 11
// output
// 1011

// input
// 5
// output
// 101


// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	printBinary(n)
// }

// func printBinary(n int) {
// 	if n == 0 {
// 		return
// 	}
// 	printBinary(n / 2)
// 	fmt.Print(n % 2)
// }


// PSEUDOCODE

// program main 

// kamus 
// 	n : integer

// algoritma
// 	input(n)
// 	printBinary(n)

// endprogram


// procedure printBinary(n : integer)

// algoritma
// 	if n <- 0 then
// 		return
// 	endif
// 	printBinary(n div 2)
// 	output(n mod 2)
// endprocedure


// Sebuah fungsi digunakan untuk menghitung total promo yang diberikan oleh swalayan ABC. 
// Promo diskon dan cashback diberikan apabila total belanja pembeli habis dibagi 100 atau habis dibagi n 
// tetapi tidak habis dibagi 100, dengan n untuk diskon adalah 4 sedangkan untuk cashback adalah 3.

// Fungsi dipanggil dalam program dengan masukan dan keluaran sebagai berikut:

// Masukan berupa sebuah bilangan bulat yang menyatakan total belanja pembeli.
// Keluaran adalah 0 bila tidak memenuhi kondisi atau 1 bila memenuhi kondisi.

// contoh 
// input
// 7
// 68900
// 66300
// 42420
// 83760
// 79610
// 66340
// 49690
// output
// 3 3

// input
// 2
// 57960
// 71130
// output
// 1 2

// input
// 5
// 93100
// 54500
// 72200
// 82900
// 61600
// output
// 1 0

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

func diskon(totalBelanja int) int {
    if totalBelanja%100 == 0 || (totalBelanja%4 == 0 && totalBelanja%100 != 0) {
        return 1
    } else {
        return 0
    }
}

func cashback(totalBelanja int) int {
    if totalBelanja%100 == 0 || (totalBelanja%3 == 0 && totalBelanja%100 != 0) {
        return 1
    } else {
        return 0
    }
}

