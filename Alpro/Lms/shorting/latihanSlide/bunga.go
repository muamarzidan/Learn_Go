package main

import "fmt"

const nMax = 1000

type arrayBunga [nMax]string

func main() {
	var tabBunga arrayBunga
	var N int

	isiArray(&tabBunga, &N)
	mengurutkan(&tabBunga, N)
	tampilArray(tabBunga, N)
}

func panjang(bunga string) int {
	return len(bunga) - 2
}

func mengurutkan(tabBunga *arrayBunga, N int) {
	var pass, idx, i int
	var temp string

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			// if panjang(tabBunga[idx]) < panjang(tabBunga[i]) { // jika menurun
			if panjang(tabBunga[idx]) > panjang(tabBunga[i]) { // jika menaik
				idx = i
			}
			i = i + 1
		}
		temp = tabBunga[pass-1]
		tabBunga[pass-1] = tabBunga[idx]
		tabBunga[idx] = temp
		pass = pass + 1
	}
}

func isiArray(tabBunga *arrayBunga, N *int) {
	fmt.Scanln(N)
	for i := 0; i < *N; i++ {
		fmt.Scanln(&tabBunga[i])
	}
}

func tampilArray(tabBunga arrayBunga, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(tabBunga[i])
	}
}



PSEUDOCODE

constant nMax : 1000

type arrayBunga array[1..nMax] of string

program main

kamus 
	tabBunga : arrayBunga
	N : integer

algoritma
	isiArray(tabBunga, N)
	mengurutkan(tabBunga, N)
	tampilArray(tabBunga, N)

endprogram


procedure mengurutkan(in/out tabBunga : arrayBunga, in N : integer)
{
	IS : variabel N dan tabBunga merupakan lemparan dari program main/utama, untuk diurutkan berdasarkan panjang string bunga
	FS : tabBunga sudah terurut berdasarkan panjang string bunga
}

kamus 
	pass, idx, i : integer
	temp : string

algoritma
	pass <- 1
	while pass <= N-1 do
		idx <- pass - 1
		i <- pass
		while i < N do
			if panjang(tabBunga[idx]) > panjang(tabBunga[i]) then
				idx <- i
			endif
			i <- i + 1
		endwhile
		temp <- tabBunga[pass-1]
		tabBunga[pass-1] <- tabBunga[idx]
		tabBunga[idx] <- temp
		pass <- pass + 1
	endwhile

endprocedure


function panjang(bunga : string) -> integer
{
	mengebalikan panjang string bunga dikurangi 2
}

kamus
	-

algoritma
	return len(bunga) - 2

endfunction


procedure isiArray(in/out tabBunga : arrayBunga, in N : integer)
{
	IS : variabel N dan tabBunga merupakan lemparan dari program main/utama, untuk diisi dengan inputan string bunga
	FS : tabBunga sudah terisi dengan inputan string bunga
}

kamus
	i : integer

algoritma
	input(N)
	for i <- 0 to N-1 do
		input(tabBunga[i])
	endfor

endprocedure


procedure tampilArray(in/out tabBunga : arrayBunga, in N : integer)	
{
	IS : variabel N dan tabBunga merupakan lemparan dari program main/utama, untuk ditampilkan
	FS : menampilkan isi tabBunga
}
kamus
	i : integer

algoritma
	for i <- 0 to N-1 do
		output(tabBunga[i])
	endfor
	
endprocedure