package main

import "fmt"

const N int = 256

func main() {
	var T [N]int
	var n int
	isiArray(&T, &n)
	reverseArray(&T, n)
	fmt.Println(isPalindrom(T, n))

}


PSEUDOCODE DARI CODE DIATAS

program main

kamus 
	const N integer : 256
	T : array [N] of integer
	n : integer

algoritma
	isiArray(T, n)
	reverseArray(T, n)
	output(isPalindrom(T, n))

endprogram