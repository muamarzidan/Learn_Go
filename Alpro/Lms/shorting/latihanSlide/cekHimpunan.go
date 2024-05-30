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
