package main
import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var nilaiAkhir tabInt

	bacaNilai(&nilaiAkhir)
	bacaNilai(&nilaiAkhir)
	bacaNilai(&nilaiAkhir)
	bacaNilai(&nilaiAkhir)
	bacaNilai(&nilaiAkhir)
	bacaNilai(&nilaiAkhir)
	cetakNilai(nilaiAkhir)
}

func bacaNilai(NA *tabInt) {
	var penampung int

	fmt.Scan(&penampung)
	if NA.n < NMAX {
		NA.info[NA.n] = penampung
		NA.n++
	}
}

func cetakNilai(NA tabInt) {
	for i := 0; i < NA.n; i++ {
		fmt.Print(NA.info[i]," ")
	}
}