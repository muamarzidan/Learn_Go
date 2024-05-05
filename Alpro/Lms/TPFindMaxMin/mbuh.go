NOMOR 1:
package main

import "fmt"

// Deklarsi konstanta 20 elemen
const NMAX int = 20

// Deklarasi tabInt dengan total NMAX elemen
type tabInt [NMAX]int

func main() {
    // Deklarasi variabel bertipe tabInt
	A := tabInt{}
	// Deklarasi banyak elemen array
	n := 0
	// Pembacaan data bilangan
	baca(&A, &n)
	// Cetak elemen array
	cetakElemen(A, n)
	// Cetak Info
	cetakInfo(A, n)
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca bilangan bulat dan memasukkan bilangan bulat positif 
		        ke dalam array A.Pembacaan dilakukan selama terbaca 
		        bilangan positif dan n < NMAX.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var input int
	for *n < NMAX {
		fmt.Scan(&input)
		if input <= 0 {
			break
		}
		A[*n] = input
		*n++
	}
}

func cetakElemen(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Elemen array: <e1> <e2> <e3> ... <en>"
	*/
	fmt.Print("Elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen maksimum dari array A dengan banyak elemen n */
	max := A[0]
	for i := 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen minimum dari array A dengan banyak elemen n */
	min := A[0]
	for i := 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Nilai maksimum, nilai minimum, dan banyaknya elemen tercetak dengan format:
			"Nilai maksimum: <max_value>
			 Nilai minimum: <min_value>
			 Banyak elemen: <n>"
	*/
	fmt.Println("Nilai maksimum:", maksimum(A, n))
	fmt.Println("Nilai minimum:", minimum(A, n))
	fmt.Println("Banyak elemen:", n)
}

// NOMOR 2:
package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	var temp pemain
	fmt.Scan(&temp.nama)
	for temp.nama != "none" && *n < NMAX {
		fmt.Scan(&temp.nomorPunggung, &temp.posisi, &temp.tinggi)
		A[*n] = temp
		*n++
		if *n < NMAX {
			fmt.Scan(&temp.nama)
		}
	}
}

func cetakPemain(A tabPemain, n int) {
	fmt.Println("Data Pemain:")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	max := A[0].tinggi
	nama := A[0].nama
	for i := 1; i < n; i++ {
		if A[i].tinggi > max {
			max = A[i].tinggi
			nama = A[i].nama
		}
	}
	fmt.Println("Pemain dengan badan tertinggi:", nama)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	min := A[0].tinggi
	nama := A[0].nama
	for i := 1; i < n; i++ {
		if A[i].tinggi < min {
			min = A[i].tinggi
			nama = A[i].nama
		}
	}
	fmt.Println("Pemain dengan badan terendah:", nama)
}