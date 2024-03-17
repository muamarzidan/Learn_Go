package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(gdc(A, B))
}

func gdc(A, B int) int {
	if B == 0 {
		return A
	} else {
		return gdc(B, A%B)
	}
}

// func gdc(A, B int) int {
// 	var c int
// 	for B > 0 {
// 		c = A % B
// 		A = B
// 		B = c
// 	}
// 	return A
// }


// func gdc(A, B int)  {
// 	if B == 0 || A == 0 {
// 		return
// 	} else {
// 		fmt.Print(gdc(B, A%B))
// 	}
// }
