package main

import "fmt"

func displayFactors(N int, N2 *int) {
	if *N2 > N {
		return
	}
	if N % *N2 == 0 {
		fmt.Print(*N2, " ")
	}
	*N2++
	displayFactors(N, N2)
}

func main() {
	var N, N2 int
	N2 = 1
	fmt.Scanln(&N)
	displayFactors(N, &N2)
	fmt.Println()
}

// PSEUDOCODE

// program main

// kamus
// 	N, N2 : integer

// algoritma
// 	N2 <- 1
// 	input(N)
// 	displayFactors(N, N2)
// 	output()

// endprogram


// procedure displayFactors(in N : integer, in/out N2 : integer)
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
// 	displayFactors(N, N2)

// endprocedure
















// package main

// import "fmt"

// func displayFactors(n int) {
// 	for i := 1; i <= n; i++ {
// 		if n%i == 0 {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }

// func main() {
// 	var N int
// 	fmt.Scanln(&N)
// 	displayFactors(N)
// 	fmt.Println()
// }

