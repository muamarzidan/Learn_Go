package main

import "fmt"

type pema struct {
	nama        string
	poin, panjangNama int
}

const MAXIMALL int = 1024

type data [MAXIMALL]pema


func main() {
    var n int
    var himpunan data
    fmt.Scanln(&n)
    isii(&himpunan, n)
    selectt(&himpunan, n)
    showw(himpunan, n)
}

func isii(himpunan *data, n int) {
    //algoritma untuk menginput data pemain dalam bentuk array 
    // gunakan variabel lokal dan input scan 
    var nama string
    var poin int
    for i := 0; i < n; i++ {
        fmt.Scanln(&nama, &poin)
        himpunan[i].nama = nama
        himpunan[i].poin = poin
        himpunan[i].panjangNama = len(nama)
    }
}

func selectt(himpunan *data, n int) {
    //algoritma selection sort secara descending
    var maxIndex int
    for i := 0; i < n-1; i++ {
        maxIndex = i
        for j := i + 1; j < n; j++ {
            if himpunan[j].panjangNama > himpunan[maxIndex].panjangNama {
                maxIndex = j
            }
        }
        himpunan[i], himpunan[maxIndex] = himpunan[maxIndex], himpunan[i]
    }
}

func showw(himpunan data, n int) {
    //algoritma untuk menampilkan data array 
    for i := 0; i < n; i++ {
        fmt.Println(himpunan[i].nama, himpunan[i].poin)
    }
}