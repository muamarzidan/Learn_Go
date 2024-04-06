package main  

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%.2f", deret(n))
}

func deret(n int) float64 {
	if n == 1 {
		return 1
	} else {
		return 1/float64(n) + deret(n-1)
	}
}