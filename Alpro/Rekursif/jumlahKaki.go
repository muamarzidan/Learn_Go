package main

import "fmt"

func main() {
	var n, result int
	fmt.Scan(&n)
	result = jumlahKaki(n)
	fmt.Println(result)
}

func jumlahKaki(n int) int {
	if n == 0 {
		return 0
	} else if n % 2 == 0 {
		return 3 + jumlahKaki(n-1)
	} else {
		return 2 + jumlahKaki(n-1)
	}
}