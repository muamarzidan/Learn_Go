package main

import "fmt"

const NMax = 1000

func main() {
	var bilangan []int
	var input int
	
	for i := 0; i < NMax; i++ {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		bilangan = append(bilangan, input)
	}
	
	jumlah := jumlahBilangan(bilangan)
	fmt.Println(jumlah)
}

func jumlahBilangan(bilangan []int) int {
	var jumlah int
	for i := 0; i < len(bilangan); i++ {
		if bilangan[i] > 0 {
			jumlah += bilangan[i]
		} else {
			jumlah = jumlah
		}
	}
	return jumlah
}