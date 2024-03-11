package main

import "fmt"

func main() {
	var a, b, c int
	var hasil string
	fmt.Scan(&a, &b, &c)
	hasil = findMax(a, b, c)
	fmt.Print(hasil)
}

func findMax(a, b, c int) string {
	var hasil string
	if a > b && a > c {
		hasil = fmt.Sprintf("%d", a)
	} 
	if b > a && b > c {
		hasil = fmt.Sprintf("%d", b)
	} 
	if c > a && c > b {
		hasil = fmt.Sprintf("%d", c)
	}
	return hasil
}