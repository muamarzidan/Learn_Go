package main

import "fmt"

func main() {
	var hariKerja, genreAksi bool
	var jamNonton int
	fmt.Scan(&hariKerja, &jamNonton, &genreAksi)
	fmt.Println(mauNonton(hariKerja, genreAksi, jamNonton))
}


func mauNonton(harikerja, genreAksi bool, jamNonton int) bool {
	if harikerja == false && jamNonton > 19 && genreAksi == true {
		return true
	}
	return false
}