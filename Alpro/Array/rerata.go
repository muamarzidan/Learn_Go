package  main

import "fmt"

const MAXINPUT int = 10

type berat [MAXINPUT]int

func main() {
	var bb berat
	var n_data int
	var rata float64

	fmt.Scan(&n_data)
	baca(&bb, n_data)
	rata = rata_berat(bb, n_data)
	cetak_rata(bb, n_data, rata)
	cetak_kurang_rata(bb, n_data, rata)
}

func baca(berat_badan *berat, n int) {
	var i int 

	if n > MAXINPUT {
		n = MAXINPUT
	}

	for i = 0; i < n; i++ {
		fmt.Scan(&berat_badan[i])
	}
}

func cetak_rata(berat_badan berat, n int, rata float64) {
	fmt.Println("Rata-rata berat badan:", rata)
}

func cetak_kurang_rata(berat_badan berat, n int, rata float64) {
	var i, total int
	for i = 0; i < n; i++ {
		if float64(berat_badan[i]) < rata {
			total++
		}
	}
	
	fmt.Println("Jumlah berat badan kurang dari rata-rata:", total)
}

func rata_berat(berat_badan berat, n int) float64 {
	var i int
	var total int = 0

	for i = 0; i < n; i++ {
		total += berat_badan[i]
	}

	return float64(total) / float64(n)
}
