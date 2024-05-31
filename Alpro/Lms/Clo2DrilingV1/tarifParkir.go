package main

import "fmt"

type kendaraan struct {
	jenis  string
	durasi int
}

func main() {
	var tKendaraan kendaraan
	fmt.Println(hitungTarif(tKendaraan))
}

func hitungTarif(tKendaraan kendaraan) int {
	var tarif, tarifMotor, tarifMobil int

	tarifMotor = 1000
	tarifMobil = 5000

	fmt.Scan(&tKendaraan.jenis, &tKendaraan.durasi)
	for (tKendaraan.jenis == "motor" || tKendaraan.jenis == "mobil") && tKendaraan.durasi >= 0 {
		if tKendaraan.jenis == "motor" {
			tarif = tarif + tarifMotor * tKendaraan.durasi
		} else {
			tarif = tarif + tarifMobil * tKendaraan.durasi
		}
		fmt.Scan(&tKendaraan.jenis, &tKendaraan.durasi)
	}

	return tarif
}
