package main

import "fmt"

func main() {
	var x, y, faktorialX, faktorialY, permutasi int

	fmt.Scan(&x, &y)

	faktorialX = Faktorial(x)
	faktorialY = Faktorial(y)

	if x >= y {
		permutasi = Permutasi(x, y)
		fmt.Print(faktorialX, faktorialY, permutasi)
	} else {
		permutasi = Permutasi(y, x)
		fmt.Print(faktorialX, faktorialY, permutasi)
	}
}

func Faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Faktorial(n-1)
}

func Permutasi(x, y int) int {
	return Faktorial(x) / Faktorial(x-y)
}

// PSEUDOCODE

// program main

// kamus
//     x, y, faktorialX, faktorialY, permutasi : integer

// algoritma
//     input(x, y)
//     faktorialX <- Faktorial(x)
//     faktorialY <- Faktorial(y)

//     if x >= y then
//         permutasi <- Permutasi(x, y)
//         output(faktorialX, faktorialY, permutasi)
//     else
//         permutasi <- Permutasi(y, x)
//         output(faktorialX, faktorialY, permutasi)
//     endif

// endprogram


// function Faktorial(n : integer) -> integer

// kamus
//     n : integer

// algoritma
//     if n == 0 then
//         return 1
//     else
//         return n * Faktorial(n-1)
//     endif

// endfunction


// function Permutasi(x, y : integer) -> integer

// kamus
//     x, y : integer

// algoritma
//     return Faktorial(x) / Faktorial(x-y)

// endfunction
