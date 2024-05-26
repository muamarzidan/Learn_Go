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


