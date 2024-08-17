package main 
import "fmt"


type pemain struct {
	nama string
	poin int
}
const MAXIMAL int = 1024
type dataPemain [MAXIMAL]pemain

func main() {
	var himpunan dataPemain
	var n int
	
	fmt.Scanln(&n)
	isiArray(&himpunan, n)
	selectionSort(&himpunan, n)
	showArray(himpunan, n)
}

func isiArray(himpunan *dataPemain, n int) {
	var nama string
	var poin int

	for i := 0; i < n; i++ {
		fmt.Scanln(&nama, &poin)
		himpunan[i].nama = nama
		himpunan[i].poin = poin
	}
}

func selectionSort(himpunan *dataPemain, n int) {
	var minIndex int
	for i := 0; i < n-1; i++ {
		minIndex = i
		for j := i + 1; j < n; j++ {
			if himpunan[j].poin < himpunan[minIndex].poin {
				minIndex = j
			}
		}
		himpunan[i], himpunan[minIndex] = himpunan[minIndex], himpunan[i]
	}
}

func showArray(himpunan dataPemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].poin)
	}
}

