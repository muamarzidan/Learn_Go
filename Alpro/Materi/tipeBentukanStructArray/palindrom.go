package main
import "fmt"


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