package main

import "fmt"

func main() {
	var letter string
	var countA, countI, countE, countU, countO int

	for letter != "."{
		fmt.Scan(&letter)
		if letter == "a" || letter == "A"{
			countA++
		} else if letter == "i" || letter == "I"{
			countI++
		} else if letter == "u" || letter == "U"{
			countU++
		} else if letter == "e" || letter == "E"{
			countE++
		} else if letter == "o" || letter == "O"{
			countO++
		}
	}
	fmt.Println("Jumlah huruf a adalah", countA)
	fmt.Println("Jumlah huruf i adalah", countI)
	fmt.Println("Jumlah huruf u adalah", countU)
	fmt.Println("Jumlah huruf e adalah", countE)
	fmt.Println("Jumlah huruf o adalah", countO)	
}

// for i := 0; i < len(letter); i++ {
// 	if letter[i:i+1] == "a" || letter[i:i+1] == "A"{
// 		countA++
// 	} else if letter[i:i+1] == "i" || letter[i:i+1] == "I"{
// 		countI++
// 	} else if letter[i:i+1] == "e" || letter[i:i+1] == "E" {
// 		countE++
// 	} else if letter[i:i+1] == "u" || letter[i:i+1] == "U"{
// 		countU++
// 	} else if letter[i:i+1] == "o" || letter[i:i+1] == "O"{
// 		countO++
// 	}
// }