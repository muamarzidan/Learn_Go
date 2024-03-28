package main

import "fmt"


type rectangle struct {
	length, width int
	color string
	property geometry
}

type geometry struct {
	area, perimeter int
}

func main() {
	var data rectangle
	isiData(&data)
	hitung(&data)
	fmt.Println(data.property.area, data.property.perimeter)
}

func isiData(persegi *rectangle) {

	fmt.Scan(&persegi.length, &persegi.width, &persegi.color)
	
}

func hitung(persegi *rectangle) {

	persegi.property.area = persegi.length * persegi.width
}


PSEUDOCODE

type rectangle <
	length, width : integer
	color : string
	property : geometry >

type geometry <
	area, perimeter : integer >

program main

kamus
	data : rectangle

algoritma
	isiData(data)
	hitung(data)
	output(data.property.area, data.property.perimeter)

endprogram


function isiData(in/out persegi : rectangle) 
{IS : data persegi telah siap pada piranti masukan (panjang, lebar, warna), untuk mengisi data persegi}
{FS : data persegi telah terisi dari data masukan}
kamus

algoritma
	input(persegi.length, persegi.width, persegi.color)

endfunction


function hitung(in/out persegi : rectangle)
{IS : terdefinisi data persegi (panjang, lebar, warna), untuk menghitung luas dan keliling persegi}
{FS : terhitung luas dan keliling persegi}
kamus

algoritma
	persegi.property.area <- persegi.length * persegi.width

endfunction

