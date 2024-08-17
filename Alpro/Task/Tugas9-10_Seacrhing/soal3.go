package main
import "fmt"


const MAX = 100

func main() {
	himpunanA := [MAX]int{11, 28, 33, 64, 95, 16, 100, 15}
	himpunanB := [MAX]int{3, 11, 7, 28, 33, 6}
	lenA := 8
	lenB := 6

	var himpunanC [MAX]int
	lenC := gabungan(himpunanA, himpunanB, lenA, lenB, &himpunanC)

	fmt.Print("Gabungan adalahn : ")
	cetakarr(himpunanC, lenC)
}

func gabungan(himpunanA, himpunanB [MAX]int, lenA, lenB int, himpunanC *[MAX]int) int {
	var totalC int
	for i := 0; i < lenA; i++ {
		himpunanC[totalC] = himpunanA[i]
		totalC++
	}
	for i := 0; i < lenB; i++ {
		if !isiarr(himpunanB[i], *himpunanC, totalC) {
			himpunanC[totalC] = himpunanB[i]
			totalC++
		}
	}
	return totalC
}

func isiarr(isi int, array [MAX]int, ukuran int) bool {
	for i := 0; i < ukuran; i++ {
		if array[i] == isi {
			return true
		}
	}
	return false
}

func cetakarr(himpunanA [MAX]int, ukuran int) {
	for i := 0; i < ukuran; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(himpunanA[i])
	}
	fmt.Println()
}



// PSEUDOCODE

// constant MAX = 100

// program main 

// kamus 
// 	himpunanA : array of integer
// 	himpunanB : array of integer
// 	lenA : integer
// 	lenB : integer
// 	himpunanC : array of integer
// 	lenC : integer

// algoritma
// 	himpunanA <- [11, 28, 33, 64, 95, 16, 100, 15]
// 	himpunanB <- [3, 11, 7, 28, 33, 6]
// 	lenA <- 8
// 	lenB <- 6

// 	himpunanC <- gabungan(himpunanA, himpunanB, lenA, lenB)
// 	cetakarr(himpunanC, lenC)

// endprogram


// function gabungan(himpunanA, himpunanB, lenA, lenB) -> integer

// kamus 
// 	totalC : integer
// 	i : integer

// algoritma
// 	totalC <- 0
// 	for i <- 0 to lenA - 1
// 		himpunanC[totalC] <- himpunanA[i]
// 		totalC <- totalC + 1
// 	endfor
// 	for i <- 0 to lenB - 1
// 		if not isiarr(himpunanB[i], himpunanC, totalC) then
// 			himpunanC[totalC] <- himpunanB[i]
// 			totalC <- totalC + 1
// 		endif	
// 	endfor
// 	return totalC
// endfunction


// function isiarr(isi, array, ukuran) -> boolean

// kamus
// 	i : integer

// algoritma
// 	for i <- 0 to ukuran - 1
// 		if array[i] = isi then
// 			return true
// 		endif
// 	endfor
// 	return false

// endfunction


// procedure cetakarr(himpunanA, ukuran)

// kamus
// 	i : integer

// algoritma
// 	for i <- 0 to ukuran - 1
// 		if i > 0 then
// 			output(", ")
// 		endif
// 		output(himpunanA[i])
// 	endfor
// endprocedure