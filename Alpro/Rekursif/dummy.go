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