package main

import "fmt"


func main() {
	var bil int
	fmt.Scanln(&bil)
	cacah(bil)
}

func cacah(bilangan int) {
	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
	F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
	for bilangan != 0 {
		fmt.Printf("%d ", bilangan%10)
		bilangan /= 10
	}
}

// func cacah(bilangan int) {
// 	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
// 	F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
// 	var digit int

// 	if bilangan == 0 {
// 		fmt.Print("0 ")
// 		return
// 	}

// 	for bilangan > 0 {
// 		digit = bilangan % 10
// 		fmt.Print(digit, " ")
// 		bilangan = bilangan / 10
// 	}
// }

// func cacah(bilangan int) {
// 	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
// 	   F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
// 	var bilPecah int
// 	for bilangan > 0 {
// 		bilangan = bilangan % 10
// 		fmt.Print(bilPecah, " ")
// 		bilangan = bilangan / 10
// 	}
// }
