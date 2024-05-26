package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(decimalToOktal(n))
}

func decimalToOktal(n int) int {
	/*I.S. Terdefinisi nilai bilangan bulat x. 
    /*F.S. fungsi mengembalikan nilai 0 jika n= 0, jika nilain>0 maka akan mengembalikan nilai oktal*/  
    if n == 0 {
		return 0
	} else {
		// Hint penjumlahan modulo 8 dengan fungsi rekursif decimalToOktal(___)*10
		return (n % 8) + decimalToOktal(n/8)*10
	}
}