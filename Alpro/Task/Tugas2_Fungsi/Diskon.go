package main
import "fmt"


func soal2() {
	var member bool
	var belanja int
	var hasil float64

	fmt.Scan(&belanja, &member)

	hasil = cekDiskon(belanja, member)
	fmt.Print(hasil)
	fmt.Print(cekDiskon(belanja, member))
}

func cekDiskon(belanja int, member bool) float64 {
	var hasil float64
	if belanja > 100000 && member {
		hasil = float64(belanja - (belanja*10/100))
	} else if belanja > 100000 && !member {
		hasil = float64(belanja - (belanja*5/100))
	} else {
		hasil = float64(belanja)
	}
	return hasil
}



// PSEUDOCODE

// program fungsiUtama
// kamus 
// 	member : boolean
// 	belanja : integer
// 	hasilDiskon : real

// algoritma
// 	input(belanja)
// 	hasilDiskon <- cekDiskon(belanja, member)
// 	output(hasilDiskon)
// endprogram

// function cekDiskon(belanja : integer, member : boolean) -> real
// kamus
// 	hasil : real 

// algoritma
// 	if belanja > 100000 and member == true then
// 		hasil <- belanja - (belanja * 10/100)
// 	else if belanja > 100000 and member == false then
// 		hasil <- belanja - (belanja * 5/100)
// 	else 
// 		hasil <- belanja
// 	endif
// 	return hasil

// endfunction