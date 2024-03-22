package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))
}

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    // Menghitung jumlah faktor
    factors := 0
    for i := 1; i <= num; i++ {
        if num%i == 0 {
            factors++
        }
    }
    // Menentukan apakah jumlah faktor genap (kecuali 1 dan dirinya sendiri)
    return factors%2 == 0
}
