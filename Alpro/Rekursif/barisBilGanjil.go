package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(rekursifGanjil(n))
}

func rekursifGanjil(n int) int {
	if n == 1 {
		return 1
	} else {
		if n%2 == 1 {
			fmt.Println(n)
		}
		return rekursifGanjil(n - 1)
	}
}
