package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s,&n)
	recPrint(n,s)
}

func recPrint(n int, s string) {
	if n == 0 {
		return
	} else {
		fmt.Println(s)
		recPrint(n-1, s)
	}
}
