package main 
import "fmt"


const nmax = 100

type tabelFrekuensi [nmax]int

func main() {
	var n int
	var t tabelFrekuensi

	fmt.Scan(&n)
	cekHimpunan(n, t)
}

func cekHimpunan(n int, t tabelFrekuensi) {
	var x int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		t[x]++
	}
	
	valid := true
	for i := 0; i < nmax; i++ {
		if t[i] > 1 {
			valid = false
			i = nmax
		}
	}
	
	if valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Tidak Valid")
	}
}



// PSEUDOCODE

// constant nmax int : 100

// type tabelFrekuensi [0...nmax-1] of integer

// program main

// kamus 
// 	n : integer
// 	t : tabelFrekuensi

// algoritma
// 	input(n)
// 	cekHimpunan(n, t)

// endprogram


// procedure cekHimpunan(in n : integer, in/out t : tabelFrekuensi)
// {
// 	IS : Menerima inputan n dan tabel t, kemudian mengecek apakah himpunan tersebut valid atau tidak
// 	FS : Menampilkan "Valid" jika himpunan tersebut valid, dan "Tidak Valid" jika himpunan tersebut tidak valid
// }

// kamus
// 	x, i : integer
// 	valid : boolean

// algoritma
// 	for i <- 0 to n-1 do
// 		input(x)
// 		t[x] <- t[x] + 1
// 	endfor

// 	valid <- true
// 	for i <- 0 to nmax-1 do
// 		if t[i] > 1 then
// 			valid <- false
// 			i <- nmax
// 		endif
// 	endfor

// 	if valid then
// 		output("Valid")
// 	else
// 		output("Tidak Valid")	
// 	endif

// endprocedure