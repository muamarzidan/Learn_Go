package main  

import "fmt"


func main() {
	var x string
	var result, index int
	index = 0
	fmt.Scanln(&x)
	result = jumlahA(x, index)
	fmt.Println(result)
}

func jumlahA(x string, index int) int {
	if index == len(x) {
		return 0
	} else {
		if x[index] == 'a' {
			return 1 + jumlahA(x, index+1)
		} else {
			return jumlahA(x, index+1)
		}
	}
}