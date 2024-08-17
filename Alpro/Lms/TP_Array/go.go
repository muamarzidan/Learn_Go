package main
import "fmt"

const NMAX int = 10
type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int

	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	var penampung int
	fmt.Scan(&penampung)

	for penampung != 0 {
		*n++
		if penampung < 0 {
			penampung = penampung * -1
		}

		if *n-1 < NMAX {
			A[*n-1] = penampung
		}

		fmt.Scan(&penampung)
	}

	if *n > NMAX {
		*n = NMAX
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	var jumNilai int

	for i := 0; i < n; i++ {
		jumNilai = jumNilai + A[i]
	}

	return jumNilai
}

func rata_rata(A tabInt, n int) float64 {
	var jumNilai int

	for i := 0; i < n; i++ {
		jumNilai = jumNilai + A[i]
	}

	return float64(jumNilai) / float64(n)
}