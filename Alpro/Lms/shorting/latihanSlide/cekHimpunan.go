package main

import "fmt"

const nMax = 37

type tHimpunan struct {
	anggota [nMax]int
	panjang int
}

func main() {
	var himp1, himp2 tHimpunan
	var n int 
	fmt.Println("Masukkan batas jumlah anggota himpunan maksimal:")
	fmt.Scan(&n)
	bacaMasukan(&himp1, n)
	bacaMasukan(&himp2, n)
	urut(&himp1)      
	urut(&himp2) 
	if sama(himp1, himp2) {
		fmt.Println("Himpunan pertama dan kedua sama = True")
	} else {
		fmt.Println("Himpunan pertama dan kedua sama = False")
	}
}

func bacaMasukan(set *tHimpunan, n int) {
	set.panjang = 0
	fmt.Println("Masukkan anggota himpunan:")
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if !ada(*set, x) {
			set.anggota[set.panjang] = x
			set.panjang++
		}
	}
}

func ada(set tHimpunan, x int) bool {
	for i := 0; i < set.panjang; i++ {
		if set.anggota[i] == x {
			return true
		}
	}
	return false
}

func urut(set *tHimpunan) {
	var pass, idx, temp, i int
	for pass = 0; pass < set.panjang-1; pass++ {
		idx = pass
		for i = pass + 1; i < set.panjang; i++ {
			if set.anggota[idx] < set.anggota[i] {
				idx = i
			}
		}
		temp = set.anggota[pass]
		set.anggota[pass] = set.anggota[idx]
		set.anggota[idx] = temp
	}
}

func sama(set1, set2 tHimpunan) bool {
	if set1.panjang != set2.panjang {
		return false
	}
	for i := 0; i < set1.panjang; i++ {
		if set1.anggota[i] != set2.anggota[i] {
			return false
		}
	}
	return true
}



PSEUDOCODE

constant nMax integer : 37

type tHimpunan struct <
	anggota array [1..nMax] of integer
	panjang int
>

program main 

kamus
	himp1, himp2 : tHimpunan
	n : integer

algoritma
	output("Masukkan batas jumlah anggota himpunan maksimal:")
	input(n)
	bacaMasukan(himp1, n)
	bacaMasukan(himp2, n)
	urut(himp1)
	urut(himp2)
	if sama(himp1, himp2) then
		output("Himpunan pertama dan kedua sama = True")
	else
		output("Himpunan pertama dan kedua sama = False")
	endif

endprogram


procedure bacaMasukan(in/out set : tHimpunan, in n : integer)
{
	IS : varibael set dan n merupakan lemparan dari program main/utama, untuk diinputkan anggota himpunan
	FS : menampilkan anggota himpunan yang sudah diinputkan
}

kamus
	i, x : integer

algoritma
	set.panjang <- 0
	output("Masukkan anggota himpunan:")
	for i <- 0 to n-1 do
		input(x)
		if not ada(set, x) then
			set.anggota[set.panjang] <- x
			set.panjang <- set.panjang + 1
		endif
	endfor

endprocedure


function ada(set : tHimpunan, x : integer) -> boolean
{
	mengebalikan nilai true jika anggota himpunan sudah ada, dan false jika anggota himpunan belum ada
}

kamus
	i : integer

algoritma
	for i <- 0 to set.panjang-1 do
		if set.anggota[i] = x then
			return true
		endif
	endfor
	return false

endfunction


procedure urut(in/out set : tHimpunan)
{
	IS : variabel set merupakan lemparan dari program main/utama, untuk diurutkan anggota himpunan
	FS : menampilkan anggota himpunan yang sudah diurutkan
}

kamus
	pass, idx, temp, i : integer

algoritma
	for pass <- 0 to set.panjang-1 do
		idx <- pass
		for i <- pass+1 to set.panjang-1 do
			if set.anggota[idx] < set.anggota[i] then
				idx <- i
			endif
		endfor
		temp <- set.anggota[pass]
		set.anggota[pass] <- set.ang
		set.anggota[idx] <- temp
	endfor

endprocedure


function sama(set1, set2 : tHimpunan) -> boolean
{
	mengebalikan nilai true jika himpunan pertama dan kedua sama, dan false jika himpunan pertama dan kedua tidak sama
}

kamus
	i : integer

algoritma
	if set1.panjang not set2.panjang then
		return false
	endif
	for i <- 0 to set1.panjang-1 do
		if set1.anggota[i] not set2.anggota[i] then
			return false
		endif
	endfor
	return true

endfunction

