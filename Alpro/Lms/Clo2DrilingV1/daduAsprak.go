package main 

import "fmt"

type SkorGame struct {
	TotalA   int
	TotalB   int
	TotalC   int
	max      int
	Pemenang byte
}

var skor SkorGame

func isGanjil(n int) bool {
	return n%2 != 0
}

func maxSkor() {
	skor.max = skor.TotalA
	skor.Pemenang = 'A'
	if skor.TotalB > skor.max {
		skor.max = skor.TotalB
		skor.Pemenang = 'B'
	} else if skor.TotalC > skor.max {
		skor.max = skor.TotalC
		skor.Pemenang = 'C'
	} 
}

func gameDadu(n int) {
	var d1, d2, d3 int

	for i := 0; i < n; i++ {
		fmt.Scan(&d1, &d2, &d3)
		if isGanjil(d1) {
			skor.TotalA += d1
		}
		if isGanjil(d2) {
			skor.TotalB += d2
		}
		if isGanjil(d3) {
			skor.TotalC += d3
		}
	}
	maxSkor()
}

func main() {
	var n int
	fmt.Scan(&n)
	gameDadu(n)
	fmt.Printf("%c %d\n", skor.Pemenang, skor.max)
}