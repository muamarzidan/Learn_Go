package main
import "fmt"


func cekFaktor(N int, N2 *int) {
	if *N2 > N {
		return
	}

	if N % *N2 == 0 {
		fmt.Print(*N2, " ")
	}

	*N2++
	cekFaktor(N, N2)
}

func soal2() {
	var N, N2 int
	N2 = 1
	fmt.Scanln(&N)
	
	cekFaktor(N, &N2)
	fmt.Println()
}



// PSEUDOCODE

// program main

// kamus
// 	N, N2 : integer

// algoritma
// 	N2 <- 1
// 	input(N)
// 	cekFaktor(N, N2)
// 	output()

// endprogram


// procedure cekFaktor(in N : integer, in/out N2 : integer)
// {IS : inisial state adalah N dan N2, nilai N adalah bilangan integer yang akan dicari faktornya, nilai N2 adalah 1}
// {FS : final state adalah N2 yang sudah bertambah 1 jika N2 lebih kecil dari N, dan N2 yang sudah mencapai N jika N2 lebih besar dari N}

// algoritma
// 	if N2 > N then
// 		return
// 	endif
// 	if N mod N2 = 0 then
// 		output(N2, " ")
// 	endif
// 	N2 <- N2 + 1
// 	cekFaktor(N, N2)

// endprocedure