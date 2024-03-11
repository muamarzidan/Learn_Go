package main

import "fmt"


func main() {
	var x, y int
	fmt.Scan(&x, &y)

	for x != y {
		mengurutkan(&x, &y)
		fmt.Scan(&x, &y)
	}
}

func mengurutkan(x, y *int) {
	if *x > *y {
		*x, *y = *y, *x
	} else {
		*x, *y = *x, *y
	}
	fmt.Println(*x, *y)
}



// PSEUDOCODE

// program main

// kamus
// 	x, y : integer

// algoritma
// 	input(x, y)
// 	while x != y do
// 		mengurutkan(x, y)
// 		input(x, y)
// 	endwhile

// endprogram


// procedure mengurutkan(in/out x : integer, in/out y : integer)

// {IS : initial state terdefinisi adalah x dan y dengan bilangan bulat positif}
// {FS : final state terdefinisi adalah x dan y dengan bilangan bulat positif}

// kamus

// algoritma
// 	if x > y then
// 		x, y <- y, x
// 	else
// 		x, y <- x, y
// 	endif
// 	output(x, y)

// endprocedure