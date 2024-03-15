package main

import "fmt"


func main() {
	var c, hasil float64
	fmt.Scan(&c)
	reamur(c, &hasil)
	fmt.Print(hasil, "R ")
	fahrenheit(c, &hasil)
	fmt.Print(hasil, "F ")
	kelvin(c, &hasil)
	fmt.Print(hasil, "K")
}

func reamur(c float64, hasil *float64) {
	*hasil = c * 0.8
}

func fahrenheit(c float64, hasil *float64) {
	*hasil = c * 1.8 + 32
}

func kelvin (c float64, hasil *float64) {
	*hasil = c + 273.15
}

//PSEUDOCODE

// program main

// kamus 
// 	c, hasil : real

// algoritma
// 	input(c)
// 	reaumur(c, hasil)
// 	output(hasil, "R ")
// 	fahrenheit(c, hasil)
// 	output(hasil, "F ")
// 	kelvin(c, hasil)
// 	output(hasil, "K")

// endprogram

// procedure reaumur(in c:real, in/out hasil:real)

// algoritma
// 	output(c * 0.8)

// endprocedure

// procedure fahrenheit(in c:real, in/out hasil:real)

// algoritma
// 	output(c * 1.8 + 32)

// endprocedure

// procedure kelvin(in c:real, in/out hasil:real)

// algoritma
// 	output(c + 273.15)

// endprocedure

