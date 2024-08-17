package main
import "fmt"


type pembeli struct {
	member bool
	harga  int
}

func main() {
	var p pembeli
	
	fmt.Scan(&p.member, &p.harga)
	hitungBiaya(p)
}

func hitungBiaya(p pembeli) {
	var total float64
	if p.member {
		if p.harga >= 400000 {
			total = float64(p.harga) - float64(p.harga)*0.2
			total = total - 50000
		} else if p.harga >= 20000 {
			total = float64(p.harga) - float64(p.harga)*0.1
			total = total - 25000
		}
	} else if p.member == false {
		if p.harga >= 300000 {
			total = float64(p.harga) - float64(p.harga)*0.2
		} else if p.harga >= 100000 {
			total = float64(p.harga) - float64(p.harga)*0.1
		}
	}
	fmt.Println(total)
}
