package main
import "fmt"


type mie struct {
	tipe    string
	pedas   bool
	topping string
}

func main() {
	var m mie
	var total int

	fmt.Scan(&m.tipe, &m.pedas)
	hitungHarga(m, &total)
	fmt.Println(total)
}

func hitungHarga(m mie, totalBayar *int) {
	var harga, iterasi int

	if m.tipe == "kwetiau" {
		harga = 8000
	} else if m.tipe == "bihun" {
		harga = 7000
	} else if m.tipe == "kuning" {
		harga = 9000
	}

	if m.pedas {
		harga += 1000
	}

	for iterasi = 0; iterasi < 5; iterasi++ {
		fmt.Scan(&m.topping)
		if m.topping == "0" {
			break
		} else if m.topping == "ayam" {
			harga += 5000
		} else if m.topping == "telur" {
			harga += 3000
		} else if m.topping == "sayur" {
			harga += 2000
		} else if m.topping == "pangsit" {
			harga += 5000
		}
	}

	*totalBayar += harga
}
