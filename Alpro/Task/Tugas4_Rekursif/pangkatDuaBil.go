package main

import "fmt"

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(pangkatDua(bil1, bil2))
}

func pangkatDua(bil, pangkat int) int {
	if bil == 0 {
		return 0
	} else if pangkat == 0 {
		return 1
	} else {
		return bil * pangkatDua(bil, pangkat-1)
	}
}


// PSEUDOCODE

// program main

// kamus
// 	bil1, bil2 : integer

// algoritma
// 	input(bil1, bil2)
// 	output(pangkatDua(bil1, bil2))

// endprogram

// function pangkatDua(bil, pangkat : integer) -> integer
// {mengembalikan bilangan hasil dari bilangan yang di pangkatkan dengan pangkat tertentu}

// kamus
// 	bil, pangkat : integer

// algoritma
// 	if bil == 0 then
// 		return 0
// 	else if pangkat == 0 then
// 		return 1
// 	else
// 		return bil * pangkatDua(bil, pangkat-1)

// endfunction
