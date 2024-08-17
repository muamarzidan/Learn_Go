package main

import "fmt"


const N int = 1000
type arrOfInt [N]int

func main() {
	var T, P arrOfInt
	var total, n int
	var length int = -1

	fmt.Scan(&total)
	for j := 0; j < total; j++ {
		fmt.Scan(&T[j])
	}

	prima(T, total, &length, &P)
	
	for n <= length && length != -1 {
		fmt.Print(P[n], " ")
		n++
	}
}

func prima(T arrOfInt, total int, length *int, P *arrOfInt) {
	var m int 
	for i := 0 ; i < total ; i++ {
		m = 0
		for k := 1; k <= T[i]; k++ {
			if T[i] % k == 0 {
				m++
			}
		}
		if m == 2 {
			*length++
			P[*length] = T[i]
		}
	}
}
