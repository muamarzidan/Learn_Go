package main

import "fmt"

const nMax = 1000

type arrayBunga [nMax]string

func main() {
	var tabBunga arrayBunga
	var N int

	isiArray(&tabBunga, &N)
	mengurutkan(&tabBunga, N)
	tampilArray(tabBunga, N)
}

func panjang(bunga string) int {
	return len(bunga) - 2
}

func mengurutkan(tabBunga *arrayBunga, N int) {
	var pass, idx, i int
	var temp string

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if panjang(tabBunga[idx]) < panjang(tabBunga[i]) {
				idx = i
			}
			i = i + 1
		}
		temp = tabBunga[pass-1]
		tabBunga[pass-1] = tabBunga[idx]
		tabBunga[idx] = temp
		pass = pass + 1
	}
}

func isiArray(tabBunga *arrayBunga, N *int) {
	fmt.Scanln(N)
	for i := 0; i < *N; i++ {
		fmt.Scanln(&tabBunga[i])
	}
}

func tampilArray(tabBunga arrayBunga, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(tabBunga[i])
	}
}