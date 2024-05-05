package main

import "fmt"

func main() {
    var x, digitFibo, total, i int
    fmt.Scan(&x)

    for i = 0; i < x; i++ {
        digitFibo = fibonacci(i)
        total = total + digitFibo
    }
    
    fmt.Println(total) 
}

func fibonacci(x int) int {
    var digit int
    if x <= 1 {
        digit = x
    } else {
        digit = fibonacci(x-1) + fibonacci(x-2)
    }
    return digit
}


// PSEUDOCODE

// program main

// kamus
// 	x, digitFibo, total, i : integer

// algoritma
// 	input(x)

// 	for i <- 0 to x-1 do
// 		digitFibo <- fibonacci(i)
// 		total <- total + digitFibo
// 		output(digitFibo, " ")
// 	endfor

// endprogram


// function fibonacci(x : integer) -> integer

// kamus
// 	digit : integer

// algoritma
// 	if x <= 1 then
// 		digit <- x
// 	else
// 		digit <- fibonacci(x-1) + fibonacci(x-2)
// 	endif
// 	return digit

// endfunction
