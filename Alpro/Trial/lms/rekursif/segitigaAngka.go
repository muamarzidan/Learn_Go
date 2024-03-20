// buatlah suatu prosedur rekursif yang menampilkan pola angka berbentuk segitiga seperti contoh diatas :

// masukan suatu bilangan bulat n dan angka

// keluaran pola angka berupa segitiga seperti pada contoh dengan angka awal sesuai dengan masukan

// contoh : angka yang digunakan hanya 0 hingga 9 jika angka lebih dari 9 maka akan kembali 0 (tips: gunakna mod 10)

// input : 2 2
// output :
//  2
// 34

// input : 5 1
// output :
//     1
//    23
//   456
//  7890
// 12345

// input : 4 5
// output :
// 	5
//    67
//   890
// 1234

// jadi pola akan terbentuk seperti segitiga siku menghdap ke kiri

package main  

import "fmt"

func main() {
	var n, angka int
	fmt.Scanln(&n, &angka)
	segitigaAngka(1, 1, n, angka)
}

func segitigaAngka(i, j, n, angka int) {
	if i <= n {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print(angka)
			angka++
			if angka > 9 {
				angka = 0
			}
		}
		fmt.Println()
		segitigaAngka(i+1, j, n, angka)
	}
}



