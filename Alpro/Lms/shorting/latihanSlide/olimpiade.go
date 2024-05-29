package main

import (
	"fmt"
)

type peserta struct {
	nama string
	g, s, b int
}

type olimpiade [100]peserta

func main() {
	var tab olimpiade
	var n int

	isiArray(&tab, &n)
	sorting(&tab, n)
	tampilArray(tab, n)
}

func isiArray(t *olimpiade, n *int) {
	fmt.Scanln(n)
	for i := 0; i < *n; i++ {
		fmt.Scanf("%s %d %d %d\n", &t[i].nama, &t[i].g, &t[i].s, &t[i].b)
	}
}

func tampilArray(t olimpiade, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(t[i].nama, t[i].g, t[i].s, t[i].b)
	}
}

func poin(g, s, b int) int {
	return (4 * g) + (3 * s) + b
}

func sorting(t *olimpiade, n int) {
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i - 1
		for j >= 0 && poin(t[j].g, t[j].s, t[j].b) < poin(temp.g, temp.s, temp.b) {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = temp
	}
}