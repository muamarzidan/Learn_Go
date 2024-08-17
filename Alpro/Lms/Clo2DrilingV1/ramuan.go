package main
import "fmt"

const NMax int = 1000
type tabT [NMax]int

func main() {
	var n, i int
	var DM, M, x, y, T tabT

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&DM[i], &M[i], &x[i], &y[i])
	}

	for i = 0; i < n; i++ {
		for DM[i] >= x[i] && M[i] >= y[i] {
			DM[i] -= x[i]
			M[i] -= y[i]
			T[i]++
		}
		fmt.Println("total ramuan yang berhasil dibuat dari lemari persediaan", i+1, "sebanyak:", T[i])
	}
}