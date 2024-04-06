package main

import (
	"fmt"
)

type Pengeluaran struct {
	namaBarang  string
	hargaBarang string
	tanggalBeli string
}

const maxData int = 100

func main() {
	var n int
	var dataPengeluaran [maxData]Pengeluaran
	var tanggalDilihat string
	//  lakukan pengisian bilangan bulat n
	fmt.Scan(&n)
	//  lakukan pengisian array dataPengeluaran sebanyak n data
	for i := 0; i < n; i++ {
		fmt.Scan(&dataPengeluaran[i].namaBarang, &dataPengeluaran[i].hargaBarang, &dataPengeluaran[i].tanggalBeli)
	}
	//  masukkan tanggal yang ingin dilihat pengeluarannya
	fmt.Scan(&tanggalDilihat)
	//  tampilkan daftar nama dan harga barang sesuai tanggal
	fmt.Println("Daftar Pengeluaran pada tanggal:", tanggalDilihat)
	for j := 0; j < n; j++ {
		if dataPengeluaran[j].tanggalBeli == tanggalDilihat {
			fmt.Println(dataPengeluaran[j].namaBarang, dataPengeluaran[j].hargaBarang)
		}
	}
}
