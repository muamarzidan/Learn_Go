package main

import "fmt"

func segitigaSiku(N int, N2 *int) {
	if *N2 > N {
		return
	}
	for j := 1; j <= *N2; j++ {
		fmt.Print("*", " ")
	}
	fmt.Println()
	*N2++
	segitigaSiku(N, N2)
}

func main() {
	var N, N2 int
	N2 = 1
	fmt.Scanln(&N)
	segitigaSiku(N, &N2)
}

// PSEUDOCODE

// program main

// kamus
// 	N : integer

// algoritma
// 	input(N)
// 	segitigaSiku(N, 1)

// endprogram

// procedure segitigaSiku(in N : integer, in/out N2 : integer)
// {IS : inisial state adalah N dan N2, nilai N adalah bilangan integer yang akan dicari faktornya, nilai N2 adalah 1}
// {FS : final state adalah N2 yang sudah bertambah 1 jika N2 lebih kecil dari N, dan N2 yang sudah mencapai N jika N2 lebih besar dari N}

// algoritma
// 	if N2 > N then
// 		return
// 	endif

// 	for j <- 1 to N2 do
// 		output("*", " ")
// 	endfor
// 	output()
// 	N2 <- N2 + 1
// 	segitigaSiku(N, N2)

// endprocedure




// func cetakPola() {
// 	var N int
// 	fmt.Scanln(&N)
// 	pola(1, N)
// }

// func pola(baris, n int) int {
// 	if baris <= n {
// 		cetakBintang(baris)
// 		fmt.Println()
// 		return pola(baris+1, n)
// 	}
// 	return 0

// }

// func cetakBintang(jumlah int) {
// 	if jumlah > 0 {
// 		fmt.Print("* ")
// 		cetakBintang(jumlah - 1)
// 	}
// }
