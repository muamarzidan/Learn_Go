package main 

import "fmt"


const N = 1000

func main() {
	var T [N]int
	var jumlah int
	fmt.Scanln(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Println(ganjilTerbesar(T, jumlah))
}

func ganjilTerbesar(T [N]int, jumlah int) int {
    //mengembalikan bilangan ganjil terbesar dalam array T dengan banyak data sebanyak jumlah
    var maxGanjil int = -1
	for i := 0; i < jumlah; i++ {
		if T[i]%2 != 0 && T[i] > maxGanjil {
			maxGanjil = T[i]
		}
	}
	return maxGanjil
}
