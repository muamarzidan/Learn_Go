package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	hitungNCari(bilangan)
}

func hitungNCari(bilangan int) {
	var i, nilaiTengah, total, digitJumlah int
	var pAwal, pTengah, pAkhir int
	digitJumlah = bilangan
	pAkhir = bilangan % 10

	for digitJumlah > 0 {
		total++
		digitJumlah = digitJumlah / 10
	}
	fmt.Println(total) 

	if total %2 == 0 {
		nilaiTengah = total / 2
	} else {
		nilaiTengah = total / 2 + 1 
	}

	i = 1
	
	for i <= total {
		if i == nilaiTengah {
			pTengah = bilangan % 10
		} else if i == total {
			pAwal = bilangan % 10
		}
		bilangan = bilangan / 10
		i++
	}

	fmt.Println(pAwal, pTengah, pAkhir)
}
