package main

import "fmt"

const nmax = 100

type tabelFrekuensi [nmax]int

func main() {
	var n int
	var t tabelFrekuensi
	fmt.Print("Masukkan maksimal jumlah anggota himpunan A : ")
	fmt.Scan(&n)
	buatIrisan(n, t)
}

func buatIrisan(n int, t tabelFrekuensi) {
	var x int
	var A tabelFrekuensi
	var B tabelFrekuensi
	var Irisan tabelFrekuensi
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		A[x]++
	}
	
	fmt.Print("Masukkan maksimal jumlah anggota himpunan B : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		B[x]++
	}
	
	for i := 0; i < nmax; i++ {
		if A[i] > 0 && B[i] > 0 {
			Irisan[i]++
		}
	}
	
	for i := 0; i < nmax; i++ {
		if Irisan[i] > 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}



PSEUDOCODE	

constant nmax int : 100

type tabelFrekuensi [nmax]int

program main

kamus
	n : integer
	t : tabelFrekuensi

algoritma
	output("Masukkan maksimal jumlah anggota himpunan A : ")
	input(n)
	buatIrisan(n, t)

endprogram
	
procedure buatIrisan(in n : integer, in/out t : tabelFrekuensi)
{
	IS : menerima inputan n dan t, kemudian setiap n dijadikan maksimal jumlah anggota himpunan A dan B, kemudian menampilkan irisan dari kedua himpunan tersebut
	FS : menampilkan irisan dari kedua himpunan tersebut
}

kamus
	x, i: integer
	A : tabelFrekuensi

algoritma
	for i <- 0 to n-1 do
		input(x)
		A[x] <- A[x] + 1
	endfor

	output("Masukkan maksimal jumlah anggota himpunan B : ")
	input(n)
	for i <- 0 to n-1 do
		input(x)
		t[x] <- t[x] + 1
	endfor

	for i <- 0 to nmax-1 do
		if A[i] > 0 and B[i] > 0 then
			t[i] <- t[i] + 1
		endif
	endfor

	for i <- 0 to nmax-1 do
		if t[i] > 0 then
			output(i, " ")
		endif
	endfor

	output()

endprocedure