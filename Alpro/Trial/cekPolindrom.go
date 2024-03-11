package main

import "fmt"

func main() {
	var kata string
	fmt.Scanf("%s", &kata)

	fmt.Print(cekKata(kata))
}

func cekKata(kata string) bool {
	var hasil string
	var cekBool bool
	var i int

	for i = len(kata) - 1; i >= 0; i-- {
		hasil = hasil + string(kata[i])
	}

	cekBool = kata == hasil
	return cekBool
}

// PSEUDOCODE

// program main

// kamus
// 	kata : string

// algoritma
// 	input(kata)
// 	output(cekKata(kata))

// endprogram

// function cekKata(kata string) bool
// // {mengembalikan boolean apakah kata merupakan polindrom atau bukan}
// kamus
// 	hasil : string
// 	cekBool : boolean
// 	i : integer

// algoritma
// 	for i <- len(kata) - 1 downto 0 do
// 		hasil <- hasil + string(kata[i])
// 	endfor

// 	cekBool <- kata = hasil
// 	return cekBool

// endfunction
