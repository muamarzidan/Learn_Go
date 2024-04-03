// Buatlah sebuah tipe bentukan nasabah digunakan untuk menyimpan data nasabah 
// (id nasabah, nama nasabah, nama bank dan nomor rekening) suatu bank di Indonesia.

// Buatlah sebuah tipe array tabNasabah dengan kapasitas 2022 orang nasabah.

// Lengkapi subprogram berikut ini sesuai dengan spesifikasi yang diberikan!

// procedure addNasabah(in/out t: tabNasabah, N : integer)
// { I.S. : terdefinisi array T yang berisi data sejumlah N nasabah }	
// { F.S. : array T bertambah 1 orang nasabah baru, tampilkan "data penuh" apabila array telah penuh }

// procedure cetak(in T:tabNasabah, N:integer, X:string) 
// { I.S. : terdefinisi array T yang berisi data sejumlah N nasabah dan string X }
// {F.S. : menampilkan data nasabah yang memiliki nama bank X}

package main

import "fmt"


type nasabah struct {
	idNasabah, namaNasabah, namaBank, noRekening string
}

type tabNasabah struct {
	data [2022]nasabah
}

func addNasabah(t *tabNasabah, N *int) {
	if *N == 2022 {
		fmt.Println("data penuh")
		return
	}
	fmt.Scan(&t.data[*N].idNasabah, &t.data[*N].namaNasabah, &t.data[*N].namaBank, &t.data[*N].noRekening)
	*N++
}

func cetak(T tabNasabah, N int, X string) {
	for i := 0; i < N; i++ {
		if T.data[i].namaBank == X {
			fmt.Println(T.data[i].idNasabah, T.data[i].namaNasabah, T.data[i].namaBank, T.data[i].noRekening)
		}
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	var T tabNasabah
	for i := 0; i < N; i++ {
		fmt.Scan(&T.data[i].idNasabah, &T.data[i].namaNasabah, &T.data[i].namaBank, &T.data[i].noRekening)
	}
	var X string
	fmt.Scan(&X)
	addNasabah(&T, &N)
	cetak(T, N, X)
}

contoh masukan
3
1234 Andi BRI 123456
5678 Budi BCA 654321
9101 Caca BNI 987654
BRI
contoh keluaran
1234 Andi BRI 123456
