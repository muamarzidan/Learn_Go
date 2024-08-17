package main
import "fmt"	


const nMax int = 100
type dataMobil [nMax]string

func main() {
	var arrMobil dataMobil
	isiData(&arrMobil)
	fmt.Print("Mobil terbanyak : ", hitung(arrMobil))
}

func isiData(arrMobil *dataMobil) {
	var i int
	for i = 0; i < nMax; i++ {
		fmt.Scanln(&arrMobil[i])
		if arrMobil[i] == "-1" {
			break
		}
	}

}

func hitung(arrMobil dataMobil) string {
	var n, nMerah, nHitam, nAbu int
	var terbanyak string

	n = len(arrMobil)
	for i := 0; i < n; i++ {
		if arrMobil[i] == "merah" {
			nMerah++
		} else if arrMobil[i] == "hitam" {
			nHitam++
		} else if arrMobil[i] == "abu" {
			nAbu++
		}
	}

	if nMerah > nHitam && nMerah > nAbu {
		terbanyak = "merah"
	} else if nHitam > nMerah && nHitam > nAbu {
		terbanyak = "hitam"
	} else {
		terbanyak = "abu"
	}

	return terbanyak
}


