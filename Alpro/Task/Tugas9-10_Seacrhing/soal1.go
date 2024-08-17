package main 
import "fmt"


func main() {
	var A [10]int
	isiArray(&A)
	fmt.Println(isValid(A))
}

func isiArray(A *[10]int) {
	for i := 0; i < 10; i++ {
		fmt.Scan(&A[i])
	}
}

func isValid(A [10]int) string {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if A[i] == A[j] {
				return "Tidak Valid"
			}
		}
	}
	return "Valid"
}

