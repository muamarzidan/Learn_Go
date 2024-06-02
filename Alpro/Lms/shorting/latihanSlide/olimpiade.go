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


PSEUDOCODE

type peserta struct <
	nama : string
	g, s, b : integer
>

type olimpiade: array[1..100] of peserta


program main

kamus 
	tab: olimpiade
	n: integer

algoritma
	isiArray(tab, n)
	sorting(tab, n)
	tampilArray(tab, n)

endprogram


procedure isiArray(in/out t : olimpiade, in n: integer)
{
	IS : variabel t bertipe olimpiade, n bertipe integer merupakan lemparan dari program main/utama
	FS : t diisi dengan data peserta olimpiade sebanyak n kali
}
kamus
	i : integer

algoritma
	input(n)
	for i <- 0 to n-1
		input(t[i].nama, t[i].g, t[i].s, t[i].b)
	endfor

endprocedure


procedure tampilArray(in/out t: olimpiade, in n: integer)
{
	IS : variabel t bertipe olimpiade, n bertipe integer merupakan lemparan dari program main/utama
	FS : menampilkan data peserta olimpiade sebanyak n kali
}
kamus
	i: integer

algoritma
	for i <- 0 to n-1
		output(t[i].nama, t[i].g, t[i].s, t[i].b)
	endfor

endprocedure


function poin(g, s, b : integer) -> integer
{
	mengembalikan nilai poin dari peserta olimpiade
}
kamus
	hasil: integer

algoritma
	hasil <- (4 * g) + (3 * s) + b
	return hasil

endfunction


procedure sorting(in/out t: olimpiade, in n: integer)
{
	IS : variabel t bertipe olimpiade, n bertipe integer merupakan lemparan dari program main/utama
	FS : mengurutkan data peserta olimpiade berdasarkan poin tertinggi menggunakan sistem algoritma selection sort
}
kamus
	pass, idx, i: integer
	temp: peserta

algoritma
	pass <- 1
	while pass <= n-1
		idx <- pass - 1
		i <- pass
		while i < n
			if poin(t[idx].g, t[idx].s, t[idx].b) < poin(t[i].g, t[i].s, t[i].b) then
				idx <- i
			endif
			i <- i + 1
		endwhile
		temp <- t[pass-1]
		t[pass-1] <- t[idx]
		t[idx] <- temp
		pass <- pass + 1
	endwhile
	
endprocedure