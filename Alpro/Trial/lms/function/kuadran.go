package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("Kuadran I")
	} else if x < 0 && y > 0 {
		fmt.Println("Kuadran II")
	} else if x < 0 && y < 0 {
		fmt.Println("Kuadran III")
	} else if x > 0 && y < 0 {
		fmt.Println("Kuadran IV")
	} else {
		fmt.Println("TITIK PUSAT")
	}
}

// func main() {
// 	var x, y float64
// 	fmt.Scan(&x, &y)

// 	result := checkQuadrant(x, y)
// 	fmt.Println(result)
// }

// func checkQuadrant(x, y float64) string {
// 	if x > 0 && y > 0 {
// 		return "Kuadran I"
// 	} else if x < 0 && y > 0 {
// 		return "Kuadran II"
// 	} else if x < 0 && y < 0 {
// 		return "Kuadran III"
// 	} else if x > 0 && y < 0 {
// 		return "Kuadran IV"
// 	} else {
// 		return "TITIK PUSAT"
// 	}
// }
