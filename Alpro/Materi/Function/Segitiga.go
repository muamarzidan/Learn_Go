package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanln(&a, &b, &c)
	fmt.Println(cekSegitiga(a, b, c))
}

func cekSegitiga(a, b, c int) string {
	if a+b > c && a+c > b && b+c > a {
		return "segitiga"
	} else {
		return "bukan segitiga"
	}
}