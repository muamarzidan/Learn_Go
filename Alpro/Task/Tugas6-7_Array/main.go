package main

import "fmt"

const N int = 256

func main() {
	var T [N]int
	var n int
	isiArray(&T, &n)
	reverseArray(&T, n)
	fmt.Println(isPalindrom(T, n))

}

func isiArray(T *[N]int, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&T[i])
	}
}

func reverseArray(T *[N]int, n int) {
	for i := 0; i < n/2; i++ {
		T[i], T[n-i-1] = T[n-i-1], T[i]
	}
}

func isPalindrom(T [N]int, n int) bool {
	for i := 0; i < n/2; i++ {
		if T[i] != T[n-i-1] {
			return false
		}
	}
	return true
}



// PSEUDOCODE

// program main

// kamus 
// 	const N integer : 256
// 	T : array [N] of integer
// 	n : integer

// algoritma
// 	isiArray(T, n)
// 	reverseArray(T, n)
// 	output(isPalindrom(T, n))

// endprogram