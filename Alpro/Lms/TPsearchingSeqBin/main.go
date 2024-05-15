package main

import "fmt"

// Deklarasi konstanta NMAX dengan nilai 10
const NMAX = 10

// Deklarasi tipe alias tabInt untuk array bilangan bulat dengan
// kapasitas maksimum NMAX
type tabInt [NMAX]int

func main() {

	// Deklarasi variabel
	var data tabInt
	var nData int
	var x1 int

	// Baca bilangan yang dicari
	fmt.Scan(&x1)

	// Baca data array
	fmt.Scan(&nData)
	bacaData(&data, &nData)

	// Cetak data
	cetakData(data, nData)

	// Pencarian bilangan tertentu dengan sequential search dan frekuensinya
	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x1) {
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.\n", frekuensiBilangan(data, nData, x1), x1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	/*
		IS: Array bilangan bulat A dan n terdefinisi sembarang
		FS: Array bilangan bulat A terisi data sebanyak n elemen.
			n tidak boleh lebih dari NMAX.
	*/
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	/*
		IS: Array bilangan A dan banyak elemen n terdefinisi
		FS: Tercetak di layar bilangan dengan format:
			"Data Bilangan: <e1> <e2> <e3> ..<en>"
	*/
	fmt.Print("Data Bilangan: ")

	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}

	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	/*
		Diberikan array bilangan A sebanyak n elemen dan bilangan x yang
		ingin dicari tahu frekuensinya, fungsi mengembalikan frekuensi
		x dalam array A.
	*/
	var counter int

	for i := 0; i < n; i++ {
		if A[i] == x {
			counter++
		}
	}

	return counter
}

func sequentialSearch(A tabInt, n int, x int) bool {
	/*
		Diberikan array bilangan A dengan banyak elemen n dan bilangan
		yang dicari (x), fungsi mengembalikan boolean true jika ditemukan,
		atau boolean false jika tidak ditemukan.
		Gunakan algoritma pencarian beruntun atau sequential search.
	*/
	var cekBilX bool = false

	for i := 0; i < n; i++ {
		if A[i] == x {
			cekBilX = true
			break
		}
	}

	return cekBilX
}










package main

import "fmt"

// Deklarasi konstanta NMAX dengan nilai 10
const NMAX = 10

// Deklarasi tipe alias tabInt untuk array bilangan bulat dengan
// kapasitas maksimum NMAX
type tabInt [NMAX]int

func main() {

	// Deklarasi variabel
	var data tabInt
	var nData int
	var x2 int
	var idx int

	// Baca bilangan yang dicari
	fmt.Scan(&x2)

	// Baca data array
	bacaData(&data, &nData)

	// Cetak data
	cetakData(data, nData)

	// Pencarian bilangan tertentu dengan binary search
	fmt.Print("Hasil pencarian: ")
	binarySearch(data, nData, x2, &idx)

	// Jika idx tidak bernilai -1, maka bilangan yang dicari ditemukan
	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	/*
		IS: Array bilangan bulat A dan n terdefinisi sembarang
		FS: Array bilangan bulat A terisi data sebanyak n elemen.
			Elemen array terisi bilangan bulat menaik (ascending).
			Jika bilangan yang dibaca lebih kecil dari
			bilangan sebelumnya, data tidak masuk ke dalam array dan
			pembacanaan terhenti.
	*/
	var temp int
	fmt.Scan(&A[0])

	*n = 1
	fmt.Scan(&temp)

	for temp >= A[*n-1] {
		if *n >= NMAX {
			*n = NMAX
			break
		}

		A[*n] = temp
		*n++
		fmt.Scan(&temp)
	}

}

func cetakData(A tabInt, n int) {
	/*
		IS: Array bilangan A dan banyak elemen n terdefinisi
		FS: Tercetak di layar bilangan dengan format:
			"Data bilangan: <e1> <e2> <e3> ..<en>"
	*/
	fmt.Print("Data Bilangan: ")

	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}

	fmt.Println()
}

func binarySearch(A tabInt, n int, x int, idx *int) {
	/*
		IS: Array bilangan A sebanyak n elemen dan bilangan x yang
		    dicari terdefinisi
		FS: idx berisi indeks lokasi x berada jika x ditemukan.
		    Jika tidak ditemukan bernilai -1.
	*/
	*idx = -1

	for i := 0; i < n; i++ {
		if A[i] == x {
			*idx = i + 1
			break
		}
	}
}