package main
import "fmt"


type nasabah struct {
	idNasabah, namaNasabah, namaBank, noRekening string
}

type tabNasabah struct {
	data [2022]nasabah
}

func addNasabah(t *tabNasabah, N *int) {
	if *N == 2022 {
		fmt.Println("data penuh")
		return
	}
	fmt.Scan(&t.data[*N].idNasabah, &t.data[*N].namaNasabah, &t.data[*N].namaBank, &t.data[*N].noRekening)
	*N++
}

func cetak(T tabNasabah, N int, X string) {
	for i := 0; i < N; i++ {
		if T.data[i].namaBank == X {
			fmt.Println(T.data[i].idNasabah, T.data[i].namaNasabah, T.data[i].namaBank, T.data[i].noRekening)
		}
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	var T tabNasabah
	for i := 0; i < N; i++ {
		fmt.Scan(&T.data[i].idNasabah, &T.data[i].namaNasabah, &T.data[i].namaBank, &T.data[i].noRekening)
	}
	var X string
	fmt.Scan(&X)
	addNasabah(&T, &N)
	cetak(T, N, X)
}