package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	rekursifGanjil(n)
}

func rekursifGanjil(n int) {
	if n <= 0 {
		return 
	}
	if n%2 == 1 {
		rekursifGanjil(n - 1)
		fmt.Print(n, " ")
	} else {
		rekursifGanjil(n - 1)
	}
}


// PSEUDOCODE

// program main

// kamus 
// 	n : integer

// algoritma
// 	input(n)
// 	rekursifGanjil(n)
// endprogram


// procedure rekursifGanjil(in n : integer)
// {IS : inisial state adalah n, nilai n adalah bilangan integer}
// {FS : -}

// algoritma
// 	if n <= 0 then
// 		return 
// 	endif
// 	if n mod 2 = 1 then
// 		rekursifGanjil(n - 1)
// 		output(n, " ")
// 	else
// 		rekursifGanjil(n - 1)
// 	endif

// endprocedure