package main

import "fmt"

func main() {
	var kar byte
	fmt.Scanf("%c", &kar)
	// fmt.Printf(string(lowToUpper(kar)))
	fmt.Printf("%c", lowToUpper(kar))
}

func lowToUpper(kar byte) byte {
	var hasil byte
	if kar >= 'a' && kar <= 'z' {
		hasil = kar - 32
	} else {
		hasil = kar
	}
	return hasil
}
