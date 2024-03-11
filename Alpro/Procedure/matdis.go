package main
import "fmt"


func main()	{
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	permutasiKombinasi(a, b, c, d)
}

func permutasiKombinasi(a, b, c, d int) {
	permutasiA := permutasi(a, c)
	kombinasiA := kombinasi(a, c)
	permutasiB := permutasi(b, d)
	kombinasiB := kombinasi(b, d)

	fmt.Println(permutasiA, kombinasiA)
	fmt.Println(permutasiB, kombinasiB)
}

func permutasi(n, r int) {
	var hasil int
	hasil = faktorial(n) / faktorial(n-r)
	
	fmt.Println(hasil)
}

func kombinasi(n, r int) {
	var hasil int
	hasil = faktorial(n) / faktorial(r) / faktorial(n-r)

	fmt.Println(hasil)
}

func faktorial(n int) {
	if n == 0 {
		fmt.Println(1)
		return
	}
	fmt.Println(n * faktorial(n-1))
}



// func main() {
// 	var a, b, c, d int
// 	fmt.Scan(&a, &b, &c, &d)

// 	permutasiKombinasi(a, b, c, d)
// }

// func permutasiKombinasi(a, b, c, d int) {
// 	permutasiA := permutasi(a, c)
// 	kombinasiA := kombinasi(a, c)
// 	permutasiB := permutasi(b, d)
// 	kombinasiB := kombinasi(b, d)

// 	fmt.Println(permutasiA, kombinasiA)
// 	fmt.Println(permutasiB, kombinasiB)
// }

// func permutasi(n, r int) int {
// 	return faktorial(n) / faktorial(n-r)
// }

// func kombinasi(n, r int) int {
// 	return faktorial(n) / faktorial(r) / faktorial(n-r)
// }

// func faktorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * faktorial(n-1)
// }