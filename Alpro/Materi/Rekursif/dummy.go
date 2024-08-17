package main
import "fmt"


func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(totalJumlah(n))
}

func totalJumlah(n int) int {
	if n ==1 {
		return 1
	} else {
		return n + totalJumlah(n-1)
	}
}