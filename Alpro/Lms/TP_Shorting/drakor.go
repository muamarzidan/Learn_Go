package main 

import "fmt"

/*
Tipe bentukan strktur tDrakor dengan atribut:
	judul(string), rating(real), episode(integer), durasi(integer),
	dan durasi(integer)
*/
type tDrakor struct {
	judul 			string
	rating 			float64
	episode, durasi int
}

// Konstanta NMAX dengan nilai 5
const NMAX int = 5

// Tipe alias tabDrakor untuk array of tDrakor dengan ukuran NMAX
type tabDrakor [NMAX]tDrakor

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

	if *n > NMAX {
		*n = NMAX
	}

	// Perulangan untuk input data drakor
	for i = 0; i < *n; i++ {
		fmt.Scanln(&(*A)[i].judul, &(*A)[i].rating, &(*A)[i].episode, &(*A)[i].durasi)
	}
}

func cetakDataDrakor(A tabDrakor, n int) {
	// Deklarasi variabel i bertipe integer
	fnt.Printf("%20s %6s %6s %6s\n", "Judul", "Rating", "Jum Ep", "Durasi")

	// Perulangan untuk mencetak data drakor
	for i := 0; i < n; i++ {
		fmt.Println("%20s %6.2f %6d %6d\n", A[i].judul, A[i].rating, A[i].episode, A[i].durasi)
	}

	fmt.Println()
}

func urutMenaik(A *tabDrakor, n int) {
	var pass, i int 
	var temp tDrakor

	pass = 1

	for pass <= n-1 {
		i == pass 
		temp = (*A)[pass]

		for i > 0 && temp.durasi < (*A)[i-1].durasi {
			(*A)[i] = (*A)[i-1]
			i--
		}

		(*A)[i] = temp
		pass++
	}
}

func urutMenurun(A *tabDrakor, n int) {
	var pass, idx, i int
	var temp tDrakor

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass

		for i < n {
			if (*A)[idx].rating < (*A)[i].rating {
				idx = i
			}

			i++
		}

		temp = (*A)[pass-1]
		(*A)[pass-1] = (*A)[idx]
		(*A)[idx] = temp
		pass++
	}
}