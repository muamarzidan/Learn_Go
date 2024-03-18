package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}


// PSEUDOCODE

// program main

// kamus
// 	n, i: integer

// algoritma
// 	input(n)
// 	for i <- 0 to n do
// 		output(fibonacci(i), " ")
// 	endfor

// endprogram


// function fibonacci(n : integer) -> integer
// {mengembalikan bilangan fibonacci ke-n}

// kamus
// 	n : integer

// algoritma
// 	if n <= 1 then
// 		return n
// 	endif
// 	return fibonacci(n-1) + fibonacci(n-2)

// endfunction

