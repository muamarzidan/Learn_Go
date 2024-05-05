package main

import "fmt"

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    var permutasiA, kombinasiA, permutasiB, kombinasiB int
    permutasi(&permutasiA, a, c)
    kombinasi(&kombinasiA, a, c)
    permutasi(&permutasiB, b, d)
    kombinasi(&kombinasiB, b, d)

    fmt.Println(permutasiA, kombinasiA)
    fmt.Println(permutasiB, kombinasiB)
}

func permutasi(hasil *int, n, r int) {
    *hasil = faktorial(n) / faktorial(n-r)
}

func kombinasi(hasil *int, n, r int) {
    *hasil = faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func faktorial(n int) int {
	for i := n-1; i > 0; i-- {
		n *= i
	}
	return n
}