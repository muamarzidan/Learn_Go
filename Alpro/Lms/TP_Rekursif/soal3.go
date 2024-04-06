package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata_rekursif(N, 1, 0, 0)
}

func cetak_rerata_rekursif(N, i, jumlah int, rata float64) {
	if i > N {
		return
	} else {
		jumlah += i
		rata = float64(jumlah) / float64(i)
		fmt.Println(jumlah, rata)
		cetak_rerata_rekursif(N, i+1, jumlah, rata)
	}
}

// func main() {
// 	var N int
// 	fmt.Scan(&N)
// 	cetak_rerata(N)
// }

// func cetak_rerata(N int) {
// 	var jumlah, i int
// 	i = 1
// 	jumlah = 0
// 	cetak_rerata_rekursif(i, N, jumlah)
// }

// func cetak_rerata_rekursif(i, N, jumlah int) {
// 	if i > N {
// 		return
// 	} else {
// 		var rata float64

// 		jumlah += i
// 		rata = float64(jumlah) / float64(i)
// 		fmt.Println(jumlah, rata)
// 		cetak_rerata_rekursif(i+1, N, jumlah)
// 	}
// }