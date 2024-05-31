package main

import "fmt"

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

// with insertion sort
func sorting(tabInt *arrayInt, N int) {
	var pass, i, temp int
	pass = 1

	for pass <= N-1 {
		temp = tabInt[pass]
		i = pass
		for i > 0 && temp < tabInt[i-1] {
			tabInt[i] = tabInt[i-1]
			i = i - 1
		}
		tabInt[i] = temp
		pass = pass + 1
	}
}

func median(tabInt *arrayInt, N int) {
	var median float64
	if N%2 == 0 {
		median = float64(tabInt[N/2]+tabInt[N/2-1]) / 2.0
		fmt.Printf("%.1f\n", median)
	} else {
		fmt.Println(tabInt[N/2])
	}
}


PSEUDOCODE

constant nMax = 1000

type arrayInt array [1..nMax] of integer

program main

kamus 
	tabInt : arrayInt
	N : integer

algoritma
	input(N)
	inputArray(tabInt, N)
	sorting(tabInt, N)
	median(tabInt, N)

endprogram


procedure inputArray(tabInt : arrayInt, N : integer)

kamus
	i : integer

algoritma
	if N > nMax then
		N = nMax
	endif

	for i = 0 to N-1 do
		input(tabInt[i])
	endfoe

endprocedure


procedure sorting(tabInt : arrayInt, N : integer)

kamus
	pass, i, temp : integer

algoritma
	pass = 1
	while pass <= N-1 do
		temp = tabInt[pass]
		i = pass
		while i > 0 and temp < tabInt[i-1] do
			tabInt[i] = tabInt[i-1]
			i = i - 1
		endwhile

		tabInt[i] = temp
		pass = pass + 1
	endwhile

endprocedure


procedure median(tabInt : arrayInt, N : integer)

kamus
	median : real

algoritma
	if N mod 2 = 0 then
		median = (tabInt[N/2] + tabInt[N/2-1]) / 2.0
		output(median)
	else
		output(tabInt[N/2])
	endif

endprocedure