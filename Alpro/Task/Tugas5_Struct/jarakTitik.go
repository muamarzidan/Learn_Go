package main

import "fmt"


type titik struct {
	x, y float64
}

func main() {
	var p1, p2 titik
	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)
	fmt.Println(jarak(p1, p2))
}

func jarak(p1, p2 titik) float64 {
	// {diterima dua buah titik, untuk mengembalikan nilai jarak antara dua titik tersebut}
	var dx, dy float64
	dx = p1.x - p2.x
	dy = p1.y - p2.y
	return akar(dx*dx + dy*dy)
}

func akar(x float64) float64 {
	// {diterima nilai x, untuk mengembalikan nilai akar dari x}
	var i int
	var akar float64
	akar = 1
	for i = 0; i < 100; i++ {
		akar = 0.5 * (akar + x/akar)
	}
	return akar
}


// PSEUDOCODE

type titik <
	x, y : real >

program main 

kamus 
	p1, p2 : titik

algoritma
	input(p1.x, p1.y, p2.x, p2.y)
	output(jarak(p1, p2))

endprogram


function jarak(p1, p2 : titik) -> real

kamus
	dx, dy : real

algoritma
	dx <- p1.x - p2.x
	dy <- p1.y - p2.y

	return akar(dx*dx + dy*dy)

endfunction


function akar(x : real) -> real

kamus
	i : integer
	akar : real

algoritma
	akar <- 1
	for i <- 0 to 100 do
		akar <- 0.5 * (akar + x/akar)
	endfor

	return akar

endfunction