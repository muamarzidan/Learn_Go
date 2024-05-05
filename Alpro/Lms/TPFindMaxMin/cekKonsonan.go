package main

import "fmt"


func serchKonsonan(x string) bool {
	var i int
	var found bool
	i = 0
	found = false

	for i < len(x) && !found {
		if x[i] != 'a' && x[i] != 'i' && x[i] != 'u' && x[i] != 'e' && x[i] != 'o' {
			found = true
		}
		i = i + 1
	}
	return found
}

func main() {
	var kata string
	fmt.Scanln(&kata)
	fmt.Println(serchKonsonan(kata))
}