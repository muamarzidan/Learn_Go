package main

import "fmt"

func NonKapital(kata string) int {
    jumlah := 0
    for i := 0; i < len(kata); i++ {
        if kata[i] >= 'a' && kata[i] <= 'z' {
            jumlah++
        }
    }
    return jumlah
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(NonKapital(input))
}