package main

import "fmt"


func main() {
	var jam, menit, detik, hasilKonversi int
	fmt.Scan(&jam, &menit, &detik)
	hasilKonversi = konversiWaktuDetik(jam, menit, detik)
	fmt.Println("Hasil konversi = ", hasilKonversi, " detik")
}

func konversiWaktuDetik(jam, menit, detik int) int {
	var hasil int
	hasil = (jam * 3600) + (menit * 60) + detik
	return hasil
}

// PSEUDOCODE

program main
kamus
    jam, menit, detik, hasilKonversi : integer

algoritma
    input(jam, menit, detik)
    hasilKonversi <- konversiWaktuDetik(jam, menit, detik)
    output("Hasil konversi = ", hasilKonversi, " detik")

endprogram


function konversiWaktuDetik(jam, menit, detik : integer) -> integer
kamus
    hasil : integer

algoritma
    hasil <- (jam * 3600) + (menit * 60) + detik
    return hasil
endfunction
