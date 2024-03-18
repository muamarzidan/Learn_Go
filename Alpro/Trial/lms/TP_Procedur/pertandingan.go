package main

import "fmt"


func main() {
	var N, g, k, jk,jm, jd, jg, jkg, jsg, jp int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&g, &k)
		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)
		hitungJumGolKegolanSelisih(g, k, &jg, &jkg, &jsg)
	}

	hitungJumPoint(jm, jd, &jp)
	fmt.Println(N, jm, jd, jk, jg, jkg, jsg, jp)
}

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg += g - k
}

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = (jm * 3) + (jd * 1)
}
