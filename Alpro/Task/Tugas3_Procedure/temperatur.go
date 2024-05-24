package main

import "fmt"

func soal4() {
	var celsius float64
	fmt.Scan(&celsius)
	reaumur(celsius)
	fmt.Print(" ")
	fahrenheit(celsius)
	fmt.Print(" ")
	kelvin(celsius)
	fmt.Println()
}

func reaumur(c float64)  {
	fmt.Print(c * 0.8)
}

func fahrenheit(c float64)  {
	fmt.Print(c * 1.8 + 32)
}

func kelvin(c float64)  {
	fmt.Print(c + 273.15)
}