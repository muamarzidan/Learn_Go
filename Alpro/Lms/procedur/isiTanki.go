package main

import "fmt"

func main() {
	var jumlah, tanki int
	fmt.Scan(&jumlah)

	isiTanki(jumlah, &tanki)
}

func isiTanki(jumlah int, tanki *int) {
	var cekPenuh bool
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&jumlah)
		if jumlah == *tanki {
			cekPenuh = true
			fmt.Println(cekPenuh)
		} else {
			fmt.Println(cekPenuh)
		}
		*tanki++
	}
}
