package main 

import "fmt"


func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(hitungJumlahDigit(n))
}

func hitungJumlahDigit(n int) int {
	if n == 0 {
		return 0
	} else {
		return 1 + hitungJumlahDigit(n/10)
	}
}


// PSEUDOCODE 

// program main 

// kamus 
// 	n : integer

// algoritma
// 	input(n)
// 	output(hitungJumlahDigit(n))

// endprogram


// function hitungJumlahDigit(n: integer) -> integer

// kamus 
// 	n : integer

// algoritma
// 	if n = 0 then
// 		return 0
// 	else
// 		return 1 + hitungJumlahDigit(n/10)

// endfunction