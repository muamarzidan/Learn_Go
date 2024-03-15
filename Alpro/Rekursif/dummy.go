package main

import "fmt"

// func main() {
// 	var a, b, n int
// 	fmt.Scan(&a, &b, &n)
// 	fmt.Println(rekursifUnartimatika(a, b, n))
// }

// func rekursifUnartimatika(a, b, n int) int {
// 	if n == 1 {
// 		return a
// 	} else {
// 		return rekursifUnartimatika(a, b, n-1) + b
// 	}
// }


func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(totalJumlah(n))
}

func totalJumlah(n int) int {
	if n ==1 {
		return 1
	} else {
		return n + totalJumlah(n-1)
	}
}

tracing
masukan 7
totalJumlah(7)
7 + totalJumlah(6)
7 + 6 + totalJumlah(5)
7 + 6 + 5 + totalJumlah(4)
7 + 6 + 5 + 4 + totalJumlah(3)
7 + 6 + 5 + 4 + 3 + totalJumlah(2)
7 + 6 + 5 + 4 + 3 + 2 + totalJumlah(1)
7 + 6 + 5 + 4 + 3 + 2 + 1
28
