package main 

import "fmt"

func main() {
	var n, m, hasil int
	fmt.Scan(&n, &m)

	hasil = perkalianLoop(n, m)
	fmt.Println(hasil)
}

func perkalianLoop(n, m int) int {
	var hasil int
	for i := 0; i < m; i++ {
		hasil = hasil + n
	}
	return hasil
}