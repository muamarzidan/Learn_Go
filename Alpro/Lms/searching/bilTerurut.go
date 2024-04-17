

package main

import "fmt"

const N int = 1000

func main() {
	var T [N]int
	var total int

	fmt.Scan(&total)
	for j := 0; j < total; j++ {
		fmt.Scan(&T[j])
	}
	fmt.Print(sudahTerurut(T, total))
}

func sudahTerurut(T [N]int, total int) bool {
    /* 
    algoritma untuk mengecek apakah data dalam array sudah terurut atau belum 
    mengembalikan boolean true jika data dalam array sudah terurut dan false jika belum    
    */
    for i := 0 ; i < total-1 ; i++ {
		if T[i] > T[i+1] {
			return false
		}
	}
	return true
}


PSEUDOCODE

constant N integer : 1000

procedure sudahTerurut (T : array [N] of integer, total : integer) -> boolean
{
	// algoritma untuk mengecek apakah data dalam array sudah terurut atau belum
	// mengembalikan boolean true jika data dalam array sudah terurut dan false jika belum
}

kamus 
	i, total : integer

algoritma
	for i <- 0 to total-1 do
		if T[i] > T[i+1] then
			return false
		endif
	endfor
	return true

endprocedure


program main 

kamus
	T : array [N] of integer
	total : integer

algoritma
	input(total)
	for j <- 0 to total-1 do
		input(T[j])
	endfor
	output(sudahTerurut(T, total))

endprogram