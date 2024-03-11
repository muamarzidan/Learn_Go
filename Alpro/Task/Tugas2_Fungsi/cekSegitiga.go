package main

import "fmt"

func main() {
	var a, b, c int
	var hasil string

	fmt.Scan(&a, &b, &c)
	hasil = cekSegitiga(a, b, c)
	fmt.Print(hasil)
}

func cekSegitiga(a, b, c int) string {
	sisia := a * 10
	sisib := b * 10
	sisic := c * 10

	if sisia + sisib + sisic == 180 {
		return "Segitiga"
	} else {
		return "Bukan Segitiga"
	}
}

// PSEUDOCODE

program main
kamus
	a, b, c : integer

algoritma
	input(a, b, c)
	output(cekSegitiga(a, b, c))

endprogram

function cekSegitiga(a, b, c : integer) -> string
kamus
	hasil : string

algoritma
	if (a + b > c) and (a + c > b) and (b + c > a) then
		hasil <- "Segitiga"
	else
		hasil <- "Bukan Segitiga"
	endif
	return hasil

endfunction
