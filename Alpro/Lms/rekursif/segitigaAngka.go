package main  
import "fmt"


func main() {
	var n, angka int
	fmt.Scanln(&n, &angka)

	segitigaAngka(1, 1, n, angka)
}

func segitigaAngka(i, j, n, angka int) {
	if i <= n {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print(angka)
			angka++
			if angka > 9 {
				angka = 0
			}
		}
		
		fmt.Println()
		segitigaAngka(i+1, j, n, angka)
	}
}



