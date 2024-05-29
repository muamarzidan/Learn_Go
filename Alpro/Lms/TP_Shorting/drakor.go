package main 

import "fmt"

/*
Tipe bentukan strktur tDrakor dengan atribut:
	judul(string), rating(real), episode(integer), durasi(integer),
	dan durasi(integer)
*/
type Drakor struct {
	judul 			string
	rating 			float64
	episode, durasi int
}

// Konstanta NMAX dengan nilai 5
const NMAX int = 5

// Tipe alias tabDrakor untuk array of tDrakor dengan ukuran NMAX
type tabDrakor [NMAX]Drakor

func main() {
	// Deklarasi variabel drakor bertipe array of tDrakor
	var drakor tabDrakor

	// Deklarasi variabel nDrakor bertipe integer untuk menampung banyaknya data drakor
	var nDrakor int

	// Pemanggilan prosedur bacaDataDrakor
	bacaDataDrakor(&drakor, &nDrakor)

	// Pemanggilan prosedur cetakDataDrakor
	fmt.Println("Data sebelum diurutkan:")
	cetakDataDrakor(drakor, nDrakor)

	// Pemanggilan prosedur urutMenurun 
	urutMenurun(&drakor, nDrakor)

	// Pemanggilan prosedur cetakDataDrakor
	fmt.Println("Data setelah diurutkan menurun berdasarkan rating:")
	cetakDataDrakor(drakor, nDrakor)

	// Pemanggilan prosedur urutNaik
	urutNaik(&drakor, nDrakor)

	// Pemanggilan prosedur cetakDataDrakor
	fmt.Println("Data setelah diurutkan menaik berdasarkan durasi:")
	cetakDataDrakor(drakor, nDrakor)
}

func bacaDataDrakor(A *tabDrakor, n *int) {
	// Deklarasi variabel i bertipe integer
	var i int

	// Input banyaknya data drakor
	fmt.Scanln(n)

	// Perulangan untuk input data drakor
	for i = 0; i < *n; i++ {
		fmt.Scanln(&A[i].judul, &A[i].rating, &A[i].episode, &A[i].durasi)
	}
}x