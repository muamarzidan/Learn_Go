package main

import "fmt"

func main() {
	var m, c float64
	fmt.Scan(&m, &c)

	fmt.Println(persamaanGaris(m, c))
}

func persamaanGaris(m, c float64) string {
	if c == 0 {
		return "melewati"
	} else {
		return "tidak melewati"
	}
}