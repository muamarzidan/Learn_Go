package main

import "fmt"


type titik struct {
	x, y float64
}

func soal2() {
	var p1, p2 titik
	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)
	fmt.Println(jarak(p1, p2))
}

func jarak(p1, p2 titik) float64 {
	var dx, dy float64
	dx = p1.x - p2.x
	dy = p1.y - p2.y
	return akar(dx*dx + dy*dy)
}

func akar(x float64) float64 {
	var i int
	var akar float64
	akar = 1
	for i = 0; i < 100; i++ {
		akar = 0.5 * (akar + x/akar)
	}
	return akar
}



// PSEUDOCODE

// program main 

// kamus 
// 	type titik <
// 	x, y : real >
// 	p1, p2 : titik

// algoritma
// 	input(p1.x, p1.y, p2.x, p2.y)
// 	output(jarak(p1, p2))

// endprogram



// function jarak(p1, p2 : titik) -> real

// kamus
// 	dx, dy : real

// algoritma
// 	dx <- p1.x - p2.x
// 	dy <- p1.y - p2.y
// 	return akar(dx*dx + dy*dy)

// endfunction



// function akar(x : real) -> real

// kamus
// 	i : integer
// 	akar : real

// algoritma
// 	akar <- 1
// 	for i <- 0 to 100 do
// 		akar <- 0.5 * (akar + x div akar)
// 	endfor
// 	return akar

// endfunction