package main

import "fmt"


Buatlah prosedur goNoGo untuk menentukan apakah seseorang akan berangkat dari rumahnya atau diam saja di rumahnya 
berdasarkan pohon keputusan dalam gambar di bawah.

              Hujan Turun?
				/		\
			tidak			ya
			/				\
		berangkat		punya payung?
						/		\
					tidak		ya	
				    /            \
			diam dirumah       berangkat

Prosedur dipanggil dalam program dengan masukan (in) dan keluaran sebagai berikut:

Masukan terdiri dari 2 buah string yang menyatakan status turunnya hujan ("ya" bila turun hujan atau "tidak" bila tidak turun hujan) 
dan status membawa ("ya" bila membawa payung atau "tidak" bila tidak membawa).
Keluaran mencetak string "berangkat" bila hujan turun atau membawa payung, atau "diam di rumah" bila tidak memenuhi kondisi tersebut.  

contoh masukan
tidak tidak
contoh keluaran
berangkat

contoh masukan
ya tidak
constoh keluaran
diam di rumah


func main() {
	var hujanTurun, bawaPayung string
	fmt.Scanln(&hujanTurun, &bawaPayung)
	goNoGo(hujanTurun, bawaPayung)
}

func goNoGo(hujanTurun, bawaPayung string) {
	if hujanTurun == "ya" {
		if bawaPayung == "ya" {
			fmt.Println("berangkat")
		} else {
			fmt.Println("diam di rumah")
		}
	} else {
		fmt.Println("berangkat")
	}
}