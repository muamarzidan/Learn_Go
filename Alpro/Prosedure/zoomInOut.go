package main

import "fmt"

func main() {
	var bilAwal, bilDua, kelipatan, hasilZoomAwal, hasilZoomDua int
	var operasi string

	fmt.Scan(&bilAwal, &bilDua, &operasi, &kelipatan)

	if operasi == "+" {
		zoomIn(bilAwal, bilDua, kelipatan, &hasilZoomAwal, &hasilZoomDua)
	} else if operasi == "-" {
		zoomOut(bilAwal, bilDua, kelipatan, &hasilZoomAwal, &hasilZoomDua)
	}
}

func zoomIn(bilAwal, bilDua, kelipatan int, hasilZoomAwal *int, hasilZoomDua *int) {
	*hasilZoomAwal = bilAwal * kelipatan
	*hasilZoomDua = bilDua * kelipatan
	fmt.Println(*hasilZoomAwal, *hasilZoomDua)
}

func zoomOut(bilAwal, bilDua, kelipatan int, hasilZoomAwal *int, hasilZoomDua *int) {
	*hasilZoomAwal = bilAwal / kelipatan
	*hasilZoomDua = bilDua / kelipatan
	fmt.Println(*hasilZoomAwal, *hasilZoomDua)
}

// PSEUDOCODE

// program main
// kamus
// 	bilAwal, bilDua, kelipatan, hasilZoomAwal, hasilZoomDua : integer
// 	operasi : string

// algoritma
// 	input(bilAwal, bilDua, operasi, kelipatan)
// 	if operasi = "+" then
// 		zoomIn(bilAwal, bilDua, kelipatan, hasilZoomAwal, hasilZoomDua)
// 	else if operasi = "-" then
// 		zoomOut(bilAwal, bilDua, kelipatan, hasilZoomAwal, hasilZoomDua)
// 	endif

// enedprogram

// procedure zoomIn(in bilAwal, bilDua, kelipatan : integer, in\out hasilZoomAwal : integer, in\out hasilZoomDua : integer)
// {IS : initial state terdefinisi adalah bilAwal, bilDua, dan kelipatan dengan bilangan bulat positif}
// {FS : final state terdefinisi adalah hasilZoomAwal dan hasilZoomDua dengan bilangan bulat positif}

// algoritma
// 	hasilZoomAwal <- bilAwal * kelipatan
// 	hasilZoomDua <- bilDua * kelipatan
// 	output(hasilZoomAwal, hasilZoomDua)

// endprocedure

// procedure zoomOut(in bilAwal, bilDua, kelipatan : integer, in\out hasilZoomAwal : integer, in\out hasilZoomDua : integer)
// {IS : initial state terdefinisi adalah bilAwal, bilDua, dan kelipatan dengan bilangan bulat positif}
// {FS : final state terdefinisi adalah hasilZoomAwal dan hasilZoomDua dengan bilangan bulat positif}

// algoritma
// 	hasilZoomAwal <- bilAwal div kelipatan
// 	hasilZoomDua <- bilDua div kelipatan
// 	output(hasilZoomAwal, hasilZoomDua)

// endprocedure
