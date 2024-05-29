package main

import "fmt"

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
	var pass, idx, i int
	var temp peserta

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if poin(t[idx].g, t[idx].s, t[idx].b) < poin(t[i].g, t[i].s, t[i].b) {
				idx = i
			}
			i = i + 1
		}
		temp = t[pass-1]
		t[pass-1] = t[idx]
		t[idx] = temp
		pass = pass + 1
	}
}