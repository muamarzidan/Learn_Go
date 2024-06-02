package main 

import "fmt"

const nmax = 100

type tHimpunan [nmax]int

func main() {
	var A tHimpunan
	var B tHimpunan
	var Gabungan tHimpunan
	var nDataA, nDataB int

	fmt.Println("Masukkan jumlah data himpunan A : ")
	fmt.Scan(&nDataA)
	fmt.Println("Masukkan jumlah data himpunan B : ")
	fmt.Scan(&nDataB)

	gabungHimpunan(nDataA, A, nDataB, B, Gabungan)
}

func gabungHimpunan(nDataA int, A tHimpunan, nDataB int, B tHimpunan, Gabungan tHimpunan) {
	var x int
	fmt.Println("Masukkan data himpunan A sebanyak", nDataA)
	for i := 0; i < nDataA; i++ {
		fmt.Scan(&x)
		A[x]++
		Gabungan[x]++
	}

	fmt.Println("Masukkan data himpunan B sebanyak", nDataB)
	for i := 0; i < nDataB; i++ {
		fmt.Scan(&x)
		B[x]++
		if Gabungan[x] == 0 {
			Gabungan[x]++
		}
	}

	for i := 0; i < nmax; i++ {
		if Gabungan[i] > 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}



PSEUDOCODE

constant nmax int : 100

type tHimpunan [nmax]int

program main

kamus
	A, B, Gabungan : tHimpunan
	nDataA, nDataB : integer

algoritma
	output("Masukkan jumlah data himpunan A : ")
	input(nDataA)
	output("Masukkan jumlah data himpunan B : ")
	input(nDataB)

	gabungHimpunan(nDataA, A, nDataB, B, Gabungan)

endprogram


procedure gabungHimpunan(in nDataA : integer, in/out A : tHimpunan, in nDataB : integer, in/out B : tHimpunan, in/out Gabungan : tHimpunan)
{
	IS : menerima inputan dari variabel nDataA dan nDataB, untuk dijadikan maksimal jumlah data ketika di inputakan himpunan A dan B
	FS : menampilkan gabungan dari himpunan A dan B dari varibael Gabungan 
}

kamus 
	x, i : integer

algoritma
	output("Masukkan data himpunan A sebanyak", nDataA)
	for i <- 0 to nDataA-1 do 
		input(x)
		A[x] <- A[x] + 1
		Gabungan[x] <- Gabungan[x] + 1
	endfor

	output("Masukkan data himpunan B sebanyak", nDataB)
	for i <- 0 to nDataB-1 do
		input(x)
		B[x] <- B[x] + 1
		if Gabungan[x] == 0 then
			Gabungan[x] <- Gabungan[x] + 1
		endif
	endfor

	for i <- 0 to nmax-1 do
		if Gabungan[i] > 0 then
			output(i, " ")
		endif
	endfor
	output()

endprocedure