// Buatlah program(dalam pseudocode) yang digunakan untuk mengurutkan seƟap dua
// bilangan yang diberikan! Implementasikan proses pengurutan dalam suatu prosedur.

// Masukan terdiri beberapa baris, yang mana seƟap barisnya terdiri dari bilangan bulat x
// dan y. Masukan akan berakhir apabila x dan y bernilai sama.
// Keluaran terdiri dari beberapa baris, yang mana masing-masing barisnya adalah nilai x dan
// y setelah diurutkan secara membesar.

// Contoh Masukan
// 3 2
// 5 1
// 4 4

// Contoh Keluaran setelah ada bilangan yang sama
// 2 3
// 1 5

// Contoh Masukan
// 30 2
// 5 1
// 4 1
// 9 2
// 3 3

// Contoh Keluaran
// 2 30
// 1 5
// 1 4
// 2 9

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	mengurutkan(&x, &y)
}

func mengurutkan(x, y *int) {
	for *x != *y {
		if *x > *y {
			*x, *y = *y, *x
		} else {
			*x, *y = *x, *y
		}
		fmt.Println(*x, *y)
		fmt.Scan(x, y)
	}
}


// Pseudocode

program main

kamus 
	x, y : integer

algoritma
	input(x, y)
	mengurutkan(x, y)

endprogram


procedure mengurutkan(in/out x, y : integerx)

kamus

algoritma
	for x != y
		if x > y then
			x, y <- y, x
		else
			x, y <- x, y
		endif
		output(x, y)
		input(x, y)
	endfor

endprocedure