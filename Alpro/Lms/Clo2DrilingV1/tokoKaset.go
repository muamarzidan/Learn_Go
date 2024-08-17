package main

import "fmt"

type KasetGame struct {
	judul      string
	harga      int
	genre      string
	statusOrder bool
}

const nmax int = 3

type tGameKaset [nmax]KasetGame

func main() {
	var kaset tGameKaset

	for i := 0; i < nmax; i++ {
		fmt.Scan(&kaset[i].judul, &kaset[i].harga, &kaset[i].genre, &kaset[i].statusOrder)
	}

	for i := 0; i < nmax; i++ {
		fmt.Println(kaset[i].judul, kaset[i].genre, hitungDiskon(kaset[i]))
	}
}

func hitungDiskon(T KasetGame) int {
	hargaSetelahDiskon := T.harga
	totalDiskon := 0

	if T.statusOrder {
		hargaSetelahDiskon -= T.harga * 5 / 100
		totalDiskon += 5
	}

	if T.statusOrder && T.harga >= 500000 {
		hargaSetelahDiskon -= T.harga * 10 / 100
		totalDiskon += 10
	}

	if T.genre == "Adventure" || T.genre == "Action" {
		hargaSetelahDiskon -= T.harga * 10 / 100
		totalDiskon += 10
	}

	if totalDiskon > 15 {
		totalDiskon = 15
		hargaSetelahDiskon = T.harga - (T.harga * 15 / 100)
	}

	return hargaSetelahDiskon
}
