package main
import "fmt"


const NMAX = 100

func main() {
    himpunanA := [NMAX]int{11, 28, 33, 64, 95, 16, 100, 15}
    himpunanB := [NMAX]int{3, 11, 7, 28, 33, 6}
    lenA := 8
    lenB := 6

    var tampung [NMAX]int
    lenC := iris(himpunanA, himpunanB, lenA, lenB, &tampung)

    fmt.Print("Irisan adalah : ")
    cetakArray(tampung, lenC)
}

func iris(himpunanA, himpunanB [NMAX]int, lenA, lenB int, tampung *[NMAX]int) int {
    var cekC int
    for i := 0; i < lenA; i++ {
        found := false
        for j := 0; j < lenB && !found; j++ {
            if himpunanA[i] == himpunanB[j] {
                if !isiKeArray(himpunanA[i], *tampung, cekC) {
                    tampung[cekC] = himpunanA[i]
                    cekC++
                }
                found = true
            }
        }
    }
    return cekC
}

func isiKeArray(isi int, array [NMAX]int, ukuran int) bool {
    for i := 0; i < ukuran; i++ {
        if array[i] == isi {
            return true
        }
    }
    return false
}

func cetakArray(A [NMAX]int, ukuran int) {
    for i := 0; i < ukuran; i++ {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(A[i])
    }
    fmt.Println()
}



// PSEUDOCODE

// constant NMAX = 100

// program main

// kamus
// 	himpunanA : array[NMAX] of integer
// 	himpunanB : array[NMAX] of integer
// 	lenA : integer
// 	lenB : integer
// 	tampung : array[NMAX] of integer
// 	lenC : integer

// algoritma
// 	himpunanA <- [11, 28, 33, 64, 95, 16, 100, 15]
// 	himpunanB <- [3, 11, 7, 28, 33, 6]
// 	lenA <- 8
// 	lenB <- 6

// 	lenC <- iris(himpunanA, himpunanB, lenA, lenB, &C)

// 	tulis("Irisan adalah : ")
// 	cetakArray(tampung, lenC)

// endprogram


// function iris(himpunanA, himpunanB : array[NMAX] of integer, lenA, lenB : integer, tampung : array[NMAX] of integer) -> integer

// kamus
// 	cekC : integer
// 	i : integer
// 	j : integer
// 	found : boolean

// algoritma
// 	cekC <- 0
// 	for i <- 0 to lenA - 1 do
// 		found <- false
// 		for j <- 0 to lenB - 1 and not found do
// 			if himpunanA[i] = himpunanB[j] then
// 				if not isiKeArray(himpunanA[i], *tampung, cekC) then
// 					tampung[nC] <- himpunanA[i]
// 					cekC <- cekC + 1
// 				endif
// 				found <- true
// 			endif
// 		endfor
// 	endfor
// 	return nC

// endfunction


// function isiKeArray(isi : integer, array : array[NMAX] of integer, ukuran : integer) -> boolean	

// kamus
// 	i : integer

// algoritma
// 	for i <- 0 to ukuran - 1 do
// 		if array[i] = isi then
// 			return true
// 		endif
// 	endfor
// 	return false

// endfunction


// procedure cetakArray(A : array[NMAX] of integer, ukuran : integer)

// kamus
// 	i : integer

// algoritma
// 	for i <- 0 to ukuran - 1 do
// 		if i > 0 then
// 			output(", ")
// 		endif
// 		output(A[i])
// 	endfor
// 	output("\n")

// endfunction