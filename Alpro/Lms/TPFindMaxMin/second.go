package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

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

func bacaData(timnas *tabPemain, nPemain *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
				Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var nama, nomorPunggung, posisi string
	var tinggi int

	fmt.Scan(&nama)
	for nama != "none" && *nPemain < NMAX {
		fmt.Scan(&nomorPunggung, &posisi, &tinggi)
		timnas[*nPemain] = pemain{nama, nomorPunggung, posisi, tinggi}
		*nPemain++
        if *nPemain < NMAX {
            fmt.Scan(&nama)
        }
	}
}

func cetakPemain(timnas tabPemain, nPemain int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	fmt.Println("Data pemain:")
	for i := 0; i < nPemain; i++ {
		fmt.Println(timnas[i].nama, timnas[i].nomorPunggung, timnas[i].posisi, timnas[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(timnas tabPemain, nPemain int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	    	"Pemain dengan badan tertingggi: <nama>"
			Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var max int = timnas[0].tinggi
	var nama string = timnas[0].nama
	for i := 1; i < nPemain; i++ {
		if timnas[i].tinggi > max {
			max = timnas[i].tinggi
			nama = timnas[i].nama
		}
	}
	fmt.Println("Pemain dengan badan tertinggi:", nama)
}

func cetakNamaPemainTerendah(timnas tabPemain, nPemain int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var min int = timnas[0].tinggi
	var nama string = timnas[0].nama
	for i := 1; i < nPemain; i++ {
		if timnas[i].tinggi < min {
			min = timnas[i].tinggi
			nama = timnas[i].nama
		}
	}
	fmt.Println("Pemain dengan badan terendah:", nama)
}
