package main

import "fmt"

func main() {
	var kapasitasBus, jumlahPenumpang int
	var cek bool
	var persen float64
	fmt.Scan(&kapasitasBus, &jumlahPenumpang)

	persen = persentase(jumlahPenumpang, kapasitasBus)

	cek = valid(persen)
	if cek {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}
}

// func persentase(jumlahPenumpang, kapasitasBus int) float64 {
// 	var hasil float64
// 	if kapasitasBus == 0 {
// 		return 0
// 	}
// 	hasil = float64(jumlahPenumpang) / float64(kapasitasBus) * 100
// 	return hasil
// }

// func valid(persentase float64) bool {
// 	return persentase >= 50 && persentase <= 75
// }

// func main() {
// 	var kapasitasBus, jumlahPenumpang int
// 	var cek bool
// 	var persen float64
// 	fmt.Scan(&kapasitasBus, &jumlahPenumpang)

// 	persen = persentase(jumlahPenumpang, kapasitasBus)

// 	cek = valid(persen)
// 	if cek {
// 		fmt.Println("berangkat")
// 	} else {
// 		fmt.Println("tidak berangkat")
// 	}
// }

// func persentase(jumlahPenumpang, kapasitasBus int) float64 {
// 	var hasil float64
// 	hasil = float64(jumlahPenumpang) / float64(kapasitasBus) * 100
// 	return hasil
// }

// func valid(persentase float64) bool {
// 	var cek bool
// 	if persentase >= 50 && persentase <= 75 {
// 		cek = true
// 	} else {
// 		cek = false
// 	}
// 	return cek
// }
