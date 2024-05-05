package main

import "fmt"

func main() {
	var banyakKonversi, i int
	var suhu float64
	fmt.Scanln(&banyakKonversi)

	i = 1

	for i <= banyakKonversi {
		fmt.Scanln(&suhu)
		f := konversiFarenhit(suhu)
		fmt.Println(suhu, "Celcius = ", f, "Fahrenheit")
		i++
	}
}

func konversiFarenhit(suhu float64) float64 {
	return (suhu * 9 / 5) + 32
}
