package main

import "fmt"

// Dalam sebuah tipe bentukan golang, terdapat 2 bentukan yaitu :
// 1. Tipe bentukan struktur / struct
// 2. Tipe bentukan array

// - Penulisan tipe bentukan struct  dalam alias : type nama_alias : tipe_data, contoh : type bilangan int, var a bilangan = 10
// - Penulisan tipe bentukan struct dalam struktur : contoh
// 	type nama_struct struct {
// 		nama_tipe_data1 : string
// 		nama_tipe_data2 : int
// 		...
// 	}

// 	pemanggilan
// 	var cobaPanggil nama_struct
// 	cobaPanggil.nama_tipe_data1 = "coba"

// - Informasi tentang array :
// 	a) Array dapat menyimpan data/elemen dengan tipe homogen/sejenis.
// 	b) Element terurut berdasarkan indeks, dimulai dari 0.
// 	c) Banyaknya elemen adalah tetap.
// 	d) Index adalah nomor untuk mengakses elemen dari array.
// 	e) Tipe indeks adalah integer atau character abjad

// - Array dibagi menjadi 2 jenis struktur:
// 	a) Array linear
// 	b) Array Non linear

// 	- Array linear : Array, list, stack, queue
// 	- Array non linear : Tree, graph

// - Penulisan array dalam golang : nama_varibel : [jumlah_elemen]tipe_data, contoh : var a [5]int



// TUGAS BUKU WITH STRUCT

type buku struct {
	title string
	author string
	page int
	price float64
}

func main() {
	buku1 := buku{
		title: "Belajar Go",
		author: "Programmer Zaman Now",	
		page: 200,
		price: 100000,
	}

	buku2 := buku{
		title: "Belajar Python",
		author: "DR Angel Ayu",
		page: 400,
		price: 300000,
	}

	buku3 := buku{
		title: "Belajar Javascript",
		author: "Maximilian Schwarzmüller",
		page: 300,
		price: 200000,
	}

	fmt.Println("Data Buku 1")
	fmt.Println("Title : ", buku1.title)
	fmt.Println("Author : ", buku1.author)
	fmt.Println("Page : ", buku1.page)
	fmt.Println("Price : ", buku1.price)
	fmt.Print("\n")
	fmt.Println("Data Buku 2")
	fmt.Println("Title : ", buku2.title)
	fmt.Println("Author : ", buku2.author)
	fmt.Println("Page : ", buku2.page)
	fmt.Println("Price : ", buku2.price)
	fmt.Print("\n")
	fmt.Println("Data Buku 3")
	fmt.Println("Title : ", buku3.title)
	fmt.Println("Author : ", buku3.author)
	fmt.Println("Page : ", buku3.page)
	fmt.Println("Price : ", buku3.price)
}


// TUGAS BUKU WITH ARRAY

type bukuArray struct {
	title    string
	author   string
	page     int
	price    float64
	kategori [4]string
}

func main() {
	buku1 := bukuArray{
		title:    "Belajar Go",
		author:   "Programmer Zaman Now",
		page:     200,
		price:    100000,
		kategori: [4]string{"Basic", "Middle", "Advance", "Expert"},
	}

	fmt.Printlxn("Data Buku :")
	fmt.Println("Title : ", buku1.title)
	fmt.Println("Author : ", buku1.author)
	fmt.Println("Page : ", buku1.page)
	fmt.Println("Price : ", buku1.price)
	for i := 0; i < len(buku1.kategori); i++ {
		fmt.Println("Kategori : ", buku1.kategori[i])
	}
}
