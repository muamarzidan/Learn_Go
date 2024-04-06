package main  
import "fmt"


type array [256]int

// Procedure untuk mengisi array dengan n bilangan
func isiArray(arr *array, n int) {
	fmt.Println("Masukkan", n, "bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

// Procedure untuk membalik isi array
func reverseArray(arr *array, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Function untuk mengecek pola palindrom pada array
func cekPalindrom(arr array, n int) bool {
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var newArr array

	fmt.Print("Masukkan jumlah bilangan yang nanti ingin di inputakan :")
	fmt.Scan(&n)

	isiArray(&newArr, n)
	fmt.Println("Isi array sebelum di reverse :", newArr)

	reverseArray(&newArr, n)
	fmt.Println("Isi array setelah di reverse :", newArr)

	if cekPalindrom(newArr, n) {
		fmt.Println("Array memiliki pola palindrom")
	} else {
		fmt.Println("Array tidak memiliki pola palindrom")
	}
}


// type array [256]int

// func isiArray(A *array, n int) {
// 	for i := 0 ; i < n; i++{
// 		fmt.Scan(&(*A)[i])
// 	}
// }

// func reverseArray(A *array, n int) {
// 	for i := 0; i < n/2; i++ {
// 		(*A)[i], (*A)[n-i-1] = (*A)[n-i-1], (*A)[i]
// 	}
// }

// func isPalindrom(A array, n int) bool {
// 	for i := 0; i < n/2; i++ {
// 		if A[i] != A[n-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	var A = array{}
// 	isiArray(&A, n)

// 	reverseArray(&A, n)
// 	fmt.Println(A)
// 	fmt.Println(isPalindrom(A, n))
// }