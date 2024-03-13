package main

import "fmt"


func main() {
	var n, m, jum int
	var rata float64
	fmt.Scanln(&n, &m)
	cetakhitungJumlahRata(n, m, &jum, &rata)
	fmt.Println(jum)
	fmt.Println(rata)
}

func cetakhitungJumlahRata(n, m int, jum *int, rata *float64) {
	var totalIterasi int
	var totalBagi int
	for i := n; i <= m; i++ {
		fmt.Println(i)
		totalIterasi = totalIterasi + i
		totalBagi++
	}
	*jum = totalIterasi
	// *rata = float64(sum) / float64(m-n+1)
	*rata = float64(totalIterasi) / float64(totalBagi)
}
