package main

import "fmt"


// func main() {
// 	var kar byte
// 	fmt.Scanf("%c", &kar)
// 	// fmt.Printf(string(lowToUpper(kar)))
// 	fmt.Printf("%c", upperToLow(kar))
// }

// func upperToLow(kar byte) byte {
// 	var hasil byte
// 	if kar >= 'A' && kar <= 'Z' {
// 		hasil = kar + 32
// 	} else {
// 		hasil = kar
// 	}
// 	return hasil
// }

func main () {
	var kata string
	fmt.Scanln(&kata)
	fmt.Println(uppernew(kata))
}

// func uppernew(kata string) string {
// 	var hasil string
// 	for i := 0; i < len(kata); i++ {
// 		hasil = hasil + string(upperToLow(kata[i]))
// 	}
// 	return hasil
// }

// func upperToLow(kar byte) byte {
// 	var hasil byte
// 	if kar >= 'A' && kar <= 'Z' {
// 		hasil = kar + 32
// 	} else {
// 		hasil = kar
// 	}
// 	return hasil
// }

func uppernew(kata string) string {
	if len(kata) == 0 {
		return ""
	}
	return string(upperToLow(kata[0])) + uppernew(kata[1:])
}

func upperToLow(kar byte) byte {
	var hasil byte
	if kar >= 'A' && kar <= 'Z' {
		hasil = kar + 32
	} else {
		hasil = kar
	}
	return hasil
}