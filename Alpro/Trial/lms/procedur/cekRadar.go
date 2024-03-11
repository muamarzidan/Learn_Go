package main

import (
	"fmt"
	"math"
)

func main() {
	var jk string
	var x, y float64
	fmt.Scanln(&jk, &x, &y)
	radar(jk, x, y)
}

func radar(jk string, x, y float64) {
	const radius = 5.0
	if jk == "tempur" && math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) <= radius {
		fmt.Println("tembak")
	} else {
		fmt.Println("tidak tembak")
	}
}

// func radar(jk string, x, y float64) {
// 	var jarak float64
// 	jarak = x*x + y*y
// 	if jk == "tempur" && jarak <= 25 {
// 		fmt.Println("tembak")
// 	} else {
// 		fmt.Println("tidak tembak")
// 	}
// }
