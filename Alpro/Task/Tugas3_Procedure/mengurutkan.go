package main

import "fmt"


func soal3() {
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

// program main

// kamus 
// 	x, y : integer

// algoritma
// 	input(x, y)
// 	mengurutkan(x, y)

// endprogram


// procedure mengurutkan(in/out x, y : integerx)

// kamus

// algoritma
// 	for x != y
// 		if x > y then
// 			x, y <- y, x
// 		else
// 			x, y <- x, y
// 		endif
// 		output(x, y)
// 		input(x, y)
// 	endfor

// endprocedure