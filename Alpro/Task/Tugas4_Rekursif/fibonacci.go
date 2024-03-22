package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

// PSEUDOCODE

function fibonacci(n : integer) -> integer

algoritma
	if n == 0 then
		return 0
	else if n == 1 then
		return 1
	else
		return fibonacci(n-1) + fibonacci(n-2)
	endif

endfunction


program main

kamus 
	n : integer

algoritma
	input(n)
	output(fibonacci(n))

endprogram
