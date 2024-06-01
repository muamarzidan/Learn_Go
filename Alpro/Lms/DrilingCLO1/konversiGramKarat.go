package main 

import "fmt"

func main() {
	var gram, karat float64
	fmt.Scan(&gram)
	karat = konversiGramKarat(gram)
	fmt.Println(karat)
}

// Buatlah fungsi saja
func konversiGramKarat(gram float64) float64 {
	// fungsi mengembalikan nilai karat dengan masukan gram
		var total float64
		total = gram * 5.00
		return total
	}