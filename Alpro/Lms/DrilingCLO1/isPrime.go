package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))
}

//buatlah functionnya saja
func isPrime(bil int) bool {
    if bil <= 1 {
        return false
    }
    factors := 0
    for i := 1; i <= bil; i++ {
        if bil%i == 0 {
            factors++
        }
    }
    return factors%2 == 0
}