package main

import "fmt"

func main() {
	var cekMember bool
	var belanja int
	fmt.Scan(&belanja, &cekMember)
	fmt.Println(cekDiskonMember(belanja, cekMember))

}

func cekDiskonMember(belanja int, cekMember bool) int {
	if belanja > 100000 && cekMember == true {
		return belanja - (belanja * 10 / 100)
	} else if belanja > 100000 && cekMember == false {
		return belanja - (belanja * 5 / 100)
	} else {
		return belanja
	}
}
