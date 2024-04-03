package main

import "fmt"

// Buatlah program dengan spesifikasi berikut ini!

// Sebuah tipe array of integer yang berkapasitas 256!
// Buatlah procedure untuk pengisian array tersebut dengan sejumlah n bilangan.
// Buatlah procedure untuk reverse isi dari array!
// Buatlah function untuk mengecek apakah suatu array memiliki pola palindrom! Nilai elemen membentuk pola simetris. 
// Contoh A = [10, 20, 30, 20, 10], B = [15, 75, 75, 15] dan C = [100]
// Buatlah program utamanya untuk menguji tiga subprogram yang telah dibuat tersebut!


func isiArray(A *[]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&(*A)[i])
	}
}

func reverseArray(A *[]int, n int) {
	for i := 0; i < n/2; i++ {
		(*A)[i], (*A)[n-i-1] = (*A)[n-i-1], (*A)[i]
	}
}

func isPalindrom(A []int, n int) bool {
	for i := 0; i < n/2; i++ {
		if A[i] != A[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	
	var A = make([]int, 256)
	isiArray(&A, n)
	reverseArray(&A, n)
	fmt.Println(A)
	fmt.Println(isPalindrom(A, n))
}
