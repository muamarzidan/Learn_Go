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
	for i := 1; i < N; i++ {
		temp := tabBunga[i]
		j := i - 1
		for j >= 0 && panjang(tabBunga[j]) > panjang(temp) {
			tabBunga[j+1] = tabBunga[j]
			j--
		}
		tabBunga[j+1] = temp
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