package main  

import "fmt"

// Buatlah prosedur dalam bahasa Go untuk menghitung bilangan terbesar dari 4 bilangan yang diinputkan.

// Masukan berupa 4 buah bilangan bulat.
// Keluaran berupa bilangan terbesar.

// contoh 
// input 
// 1 2 3 4
// output
// 4

// input
// 4 3 2 1
// output
// 4

// input
// 1 1 1 1
// output
// 1


func main() {
	var b1, b2, b3, b4 int
	fmt.Scan(&b1, &b2, &b3, &b4)
	cetakTerbesar(b1, b2, b3, b4)
}

func cetakTerbesar(b1 int, b2 int, b3 int, b4 int) {
	/*
	I.S menerima empat bilangan bulat b1, b2, b3, dan b4.
	F.S mencari dan mencetak bilangan terbesar di antara b1, b2, b3, dan b4.
	*/
	if b1 > b2 && b1 > b3 && b1 > b4 {
		fmt.Println(b1)
	} else if b2 > b1 && b2 > b3 && b2 > b4 {
		fmt.Println(b2)
	} else if b3 > b1 && b3 > b2 && b3 > b4 {
		fmt.Println(b3)
	} else {
		fmt.Println(b4)
	}
}