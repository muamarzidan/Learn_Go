package main

import "fmt"

const nMax = 37
const N = 5

type tHimpunan struct {
	anggota [nMax]int
	panjang int
}

func main() {
	var himp1, himp2 tHimpunan
	var n int 
	fmt.Println("masukan batas jumlah anggota himpunan maximal")
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
	set.panjang = n
	fmt.Println("Masukkan anggota himpunan:")
	for i := 1; i <= n; i++ {
		fmt.Scan(&set.anggota[i])
	}
}

func ada(set tHimpunan, x int) bool {
	for i := 1; i <= set.panjang; i++ {
		if set.anggota[i] == x {
			return true
		}
	}
	return false
}

func urut(set *tHimpunan) {
	for i := 2; i <= set.panjang; i++ {
		temp := set.anggota[i]
		j := i - 1
		for j >= 1 && set.anggota[j] > temp {
			set.anggota[j+1] = set.anggota[j]
			j--
		}
		set.anggota[j+1] = temp
	}
}

func sama(set1, set2 tHimpunan) bool {
	if set1.panjang != set2.panjang {
		return false
	}
	for i := 1; i <= set1.panjang; i++ {
		if set1.anggota[i] != set2.anggota[i] {
			return false
		}
	}
	return true
}
