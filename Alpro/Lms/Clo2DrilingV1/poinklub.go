package main
import "fmt"


func main() {
	type klub struct {
		nama string
		point, selisihGol int
	}
	const nmax int = 3
	var klubs [nmax]klub

	for i := 0; i < nmax; i++ {
		fmt.Scan(&klubs[i].nama, &klubs[i].point, &klubs[i].selisihGol)
	}

	var jumlahPoint int
	
	for i := 0; i < nmax; i++ {
		jumlahPoint = jumlahPoint + klubs[i].point
	}

	rataPoint := float64(jumlahPoint) / float64(nmax)
	fmt.Println(rataPoint)
}