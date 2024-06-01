package main 

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cacah(bilangan)
}

// Buatlah prosedur saja
func cacah(bilangan int) {
	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
	F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
	var digit int

	if bilangan == 0 {
		fmt.Print("0 ")
		return
	}

	for bilangan > 0 {
		digit = bilangan % 10
		fmt.Print(digit, " ")
		bilangan = bilangan / 10
	}
}