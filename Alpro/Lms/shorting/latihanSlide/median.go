package main

import (
	"fmt"
)

const nMax int = 1000

type arrayInt [nMax]int

func main() {
	var tabInt arrayInt
	var N int

	fmt.Scanln(&N)
	inputArray(&tabInt, N)

	sorting(&tabInt, N)
	median(&tabInt, N)
}

func inputArray(tabInt *arrayInt, N int) {
	if N > nMax {
		N = nMax
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&tabInt[i])
	}
}

func sorting(tabInt *arrayInt, N int) {
	for i := 1; i < N; i++ {
		temp := tabInt[i]
		j := i - 1
		for j >= 0 && tabInt[j] > temp {
			tabInt[j+1] = tabInt[j]
			j--
		}
		tabInt[j+1] = temp
	}
}

func median(tabInt *arrayInt, N int) {
	var media float64
	if N%2 == 0 {
		median = float64(tabInt[N/2]+tabInt[N/2-1]) / 2.0
		fmt.Printf("%.1f\n", median)
	} else {
		fmt.Println(tabInt[N/2])
	}
}