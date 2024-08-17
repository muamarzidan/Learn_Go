package main
import "fmt"

type barang struct {
	nama               string
	harga, panjangNama int
}

const NMAX int = 1024

type dataBarang [NMAX]barang

func main() {
	var himpunan dataBarang
	var n int
	fmt.Scanln(&n)
	isiArray(&himpunan, n)
	insertionSort(&himpunan, n)
	showArray(himpunan, n)
}

func isiArray(himpunan *dataBarang, n int) {
	//algoritma untuk menginput data himpunan barang ke dalam bentuk array
	for i := 0; i < n; i++ {
		// gunakan fmt.Scan untuk menginput nilai di sini
		fmt.Scanln(&himpunan[i].nama, &himpunan[i].harga)
		//prosem menghimpun nilai sesuai jumlah n
		himpunan[i].panjangNama = len(himpunan[i].nama)
	}
}

func insertionSort(himpunan *dataBarang, n int) {
	//algoritma insertion sort secara ascending
	for i := 1; i < n; i++ {
		temp := himpunan[i]
		j := i - 1
		for j >= 0 && himpunan[j].panjangNama > temp.panjangNama {
			himpunan[j+1] = himpunan[j]
			j--
		}
		himpunan[j+1] = temp
	}
}

func showArray(himpunan dataBarang, n int) {
	//algoritma untuk menampilkan data barang menggunakan pengulangan array
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].harga)
	}
}
