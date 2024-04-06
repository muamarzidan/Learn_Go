// ada soal golang 

// Menjelang Hari Raya Lebaran 2024 ini sebuah Software House akan membagikan bonus kepada pegawainya.
//  Besarnya bonus akan ditentukan berdasarkan gaji pokok dan angka pengali. Jika masa kerja minimal 10 tahun, 
//  maka angka pengalinya 1.5. Jika masa kerjanya di bawah 10 tahun hingga minimal 5 tahun, maka angka pengalinya 0.75. 
//  Masa kerja di bawah 5 tahun, angka pengalinya 0.5. 

// Buatlah program untuk menghitung besarnya bonus yang diterima pegawai. Definisikan sebuah tipe bentukan struct untuk 
// pegawai dengan field nama, gaji pokok, masa kerja, dan besar bonus.  Masukan dan keluaran program adalah sebagai berikut:

// Masukan berupa nama, besar gaji pokok, dan lama masa kerja. 

// Keluaran berupa teks dengan format "Pegawai dengan nama <nama> mendapatkan bonus sebesar Rp <besar_bonus>"

// For example:

// Input	Result
// Manggala 10000000 17
// Pegawai dengan nama Manggala mendapatkan bonus sebesar Rp 15000000
// Gamma 10000000 9
// Pegawai dengan nama Gamma mendapatkan bonus sebesar Rp 7500000
// Selly 10000000 5
// Pegawai dengan nama Selly mendapatkan bonus sebesar Rp 7500000
// Hanifan 10000000 4
// Pegawai dengan nama Hanifan mendapatkan bonus sebesar Rp 5000000
// Prasti 10000000 10
// Pegawai dengan nama Prasti mendapatkan bonus sebesar Rp 15000000



package main

import "fmt"

// Tipe bentukan pegawai dengan komponen (field) nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama string
	gaji_pokok, masa_kerja int
	besar_bonus float64
}

func main() {
    // deklarasi variabel bertipe pegawai
	var p pegawai
	
	// baca data pengawai
	fmt.Scan(&p.nama, &p.gaji_pokok, &p.masa_kerja)
	
	// hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&p)
	// Cetak besar bonus
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %.0f\n", p.nama, p.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	/* IS: p.nama, p.gaji_pokok, p.masa_kerja terdefinisi
	   Proses: Besar bonus dihitung dengan mengalikan masa kerja dengan angka pengali
	           Jika masa kerja minimal 10 tahun, angka pengalinya 1.5
	           jika masa kerja di bawah 10 tahun hingga 5 tahun, angka pengalinya 0.75
			   di bawah 5 tahun, angka pengalinya 0.5
	   FS: p.besar_bonus berisi nilai
	*/
	if p.masa_kerja >= 10 {
		p.besar_bonus = float64(p.gaji_pokok) * 1.5
	} else if p.masa_kerja >= 5 {
		p.besar_bonus = float64(p.gaji_pokok) * 0.75
	} else {
		p.besar_bonus = float64(p.gaji_pokok) * 0.5
	}
}
