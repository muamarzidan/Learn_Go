package main

import "fmt"

func main() {
	var n, i, j int
	fmt.Scan(&n)

	for i = 0; i < n; i++ { // i cetak baris
		for j = 0; j < n; j++ { // j cetak kolom
			if i == j {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
