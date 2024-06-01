package main 

import "fmt"

func main() {
	func mauNonton(harikerja, genreAksi bool, jamNonton int) bool {
		if harikerja == false && jamNonton > 19 && genreAksi == true {
			return true
		}
		return false
	}
}

func mauNonton(harikerja, genreAksi bool, jamNonton int) bool {
	if harikerja == false && jamNonton > 19 && genreAksi == true {
		return true
	}
	return false
}