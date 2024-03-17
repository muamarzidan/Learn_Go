package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(nilaiSukuKeN(N))
}	

func nilaiSukuKeN(N int) int {
	if N == 1 || N == 2 || N == 3 {
		return 1
	} else {
		return nilaiSukuKeN(N-1) + nilaiSukuKeN(N-2) + nilaiSukuKeN(N-3)
	}
}

tracing code diatas
misalkan N  = 10, YANG DIMANA HARUSNYA BEROUPUT 105

nilaiSukuKeN(10) = nilaiSukuKeN(9) + nilaiSukuKeN(8) + nilaiSukuKeN(7)
nilaiSukuKeN(9) = nilaiSukuKeN(8) + nilaiSukuKeN(7) + nilaiSukuKeN(6)
nilaiSukuKeN(8) = nilaiSukuKeN(7) + nilaiSukuKeN(6) + nilaiSukuKeN(5)
nilaiSukuKeN(7) = nilaiSukuKeN(6) + nilaiSukuKeN(5) + nilaiSukuKeN(4)
nilaiSukuKeN(6) = nilaiSukuKeN(5) + nilaiSukuKeN(4) + nilaiSukuKeN(3)
nilaiSukuKeN(5) = nilaiSukuKeN(4) + nilaiSukuKeN(3) + nilaiSukuKeN(2)
nilaiSukuKeN(4) = nilaiSukuKeN(3) + nilaiSukuKeN(2) + nilaiSukuKeN(1)
nilaiSukuKeN(3) = 1
nilaiSukuKeN(2) = 1
nilaiSukuKeN(1) = 1
nilaiSukuKeN(4) = 1 + 1 + 1 = 3
nilaiSukuKeN(5) = 1 + 1 + 3 = 5
nilaiSukuKeN(6) = 1 + 3 + 5 = 9
nilaiSukuKeN(7) = 3 + 5 + 9 = 17
nilaiSukuKeN(8) = 5 + 9 + 17 = 31
nilaiSukuKeN(9) = 9 + 17 + 31 = 57
nilaiSukuKeN(10) = 17 + 31 + 57 = 105