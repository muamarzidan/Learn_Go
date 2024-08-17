package main
import "fmt"


const NMAX int = 1024
type mahasiswa struct {
	nama, indeks string
	nilai  int
}
type dataMahasiswa [NMAX]mahasiswa

func main() {
	var himpunan dataMahasiswa
	var n int
	var tampilIndeks string

	fmt.Scanln(&n)
	fmt.Scanln(&tampilIndeks)
	isiArray(&himpunan, n)
	insertionSort(&himpunan, n)
	showArray(himpunan, n, tampilIndeks)
}

func isiArray(himpunan *dataMahasiswa, n int) {
	for i := 0; i < n; i++ {
		fmt.Scanln(&himpunan[i].nama, &himpunan[i].nilai)
		if himpunan[i].nilai <= 40 {
			himpunan[i].indeks = "E"
		} else if himpunan[i].nilai > 40 && himpunan[i].nilai <= 50 {
			himpunan[i].indeks = "D"
		} else if himpunan[i].nilai > 50 && himpunan[i].nilai <= 60 {
			himpunan[i].indeks = "C"
		} else if himpunan[i].nilai > 60 && himpunan[i].nilai <= 65 {
			himpunan[i].indeks = "BC"
		} else if himpunan[i].nilai > 65 && himpunan[i].nilai <= 70 {
			himpunan[i].indeks = "B"
		} else if himpunan[i].nilai > 70 && himpunan[i].nilai <= 80 {
			himpunan[i].indeks = "AB"
		} else if himpunan[i].nilai > 80 && himpunan[i].nilai <= 100 {
			himpunan[i].indeks = "A"
		}
	}
}

func insertionSort(himpunan *dataMahasiswa, n int) {
	for i := 1; i < n; i++ {
		temp := himpunan[i]
		j := i - 1
		for j >= 0 && himpunan[j].nilai > temp.nilai {
			himpunan[j+1] = himpunan[j]
			j--
		}
		himpunan[j+1] = temp
	}
}

func showArray(himpunan dataMahasiswa, n int, tampilIndeks string) {
	for i := 0; i < n; i++ {
		if himpunan[i].indeks == tampilIndeks{
			fmt.Println(himpunan[i].nama, himpunan[i].nilai, himpunan[i].indeks)
		}
	}
}

