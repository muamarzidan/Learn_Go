package main

import "fmt"

// Buatlah prosedur cetakhitungJumlahRata untuk mencetak, menjumlahkan, dan menghitung rata-rata bilangan bulat dari n hingga m, dengan n < m.
// Pada prosedur terdapat empat parameter, yaitu 2 parameter input berupa bilangan bulat n dan m dan 2 parameter berjenis input/output berupa bilangan bulat dan
// bilangan real untuk menampung hasil penjumlahan dan rata-rata sebagaimana dapat dilihat pada spesifikasi berikut

// procedure cetakhitungJumlahRata(in n, m : integer,  in/out jum : integer, rata : real)
// { I.S.: n dan m terdefinisi
//   F.S.: jum berisi nilai hasil penjumlahan bilangan bulat dari n hingga m, rata berisi nilai rata-ratanya }
// Prosedur dipanggil dalam program dengan masukan dan keluaran sebagai berikut:

// Masukan berupa dua bilangan bulat n dan m.
// Keluaran adalah hasil penjumlahan n hingga m dan rata-ratanya.

// contoh masukan
// 4 5
// contoh keluaran
// 4
// 5
// 9 4.5

// contoh masukan
// 5 10
// contoh keluaran
// 5
// 6
// 7
// 8
// 9
// 10
// 45 7.5

// contoh masukan
// -1 1
// contoh keluaran
// -1
// 0
// 1
// 0 0.0

func main() {
	var n, m, jum int
	var rata float64
	fmt.Scanln(&n, &m)
	cetakhitungJumlahRata(n, m, &jum, &rata)
	fmt.Println(jum)
	fmt.Println(rata)
}

func cetakhitungJumlahRata(n, m int, jum *int, rata *float64) {
	var totalIterasi int
	var totalBagi int
	for i := n; i <= m; i++ {
		fmt.Println(i)
		totalIterasi = totalIterasi + i
		totalBagi++
	}
	*jum = totalIterasi
	// *rata = float64(sum) / float64(m-n+1)
	*rata = float64(totalIterasi) / float64(totalBagi)
}
