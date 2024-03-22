package main

import "fmt"


// Seorang mahasiswa Tel-U bermain game 3 dadu. Dadu yang digunakan adalah dadu seimbang bermata 6. 
// Secara berulang-ulang dia melemparkan ketiga dadu itu dan mencatat kemunculan angkanya. 
// Setiap muncul angka-angka genap dari ketiga dadu itu dicatat. 
// Permainan berakhir ketika muncul jumlah angka 12 dari ketiga dadu itu. 
// Jumlah angka 12  yang dibentuk dari 3 bilangan genap tidak akan dimasukkan ke dalam perhitungan.

// Buatlah procedure dalam bahasa Go bernama hitungTigaGenap untuk membantu menghitung banyaknya kemunculan angka ketiganya genap di ketiga dadu.

// procedure hitungTigaGenap(in d1 : integer, d2: integer, d3 : integer, in/out total : integer){
// 	IS : Terdefinisi ketiga angka dadu
// 	FS : menambah total sebanyak 1 jika ketiga angka dadu genap
// }

// Masukan terdiri dari beberapa baris, dengan masing-masing baris adalah angka yang muncul pada ketiga dadu. Masukan berhenti ketika muncul jumlah 12 dari ketiga dadu itu. 

// Keluaran berupa jumlah kemunculan mata dadu ketiganya genap. 

func main() {
	var d1, d2, d3, total int
    fmt.Scan(&d1, &d2, &d3)
    total = 0
    for (d1 + d2 + d3) != 12{
        hitungTigaGenap(d1,d2,d3, &total)
        fmt.Scan(&d1, &d2, &d3)
    }
    fmt.Println(total)
}

func hitungTigaGenap(d1 int, d2 int, d3 int, total *int){
	if d1 % 2 == 0 && d2 % 2 == 0 && d3 % 2 == 0{
		*total++
	}
}
