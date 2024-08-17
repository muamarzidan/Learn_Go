package main 
import "fmt"


func main() {
	var banyaknyaPembeli, totalBelanja, i int
	var DapatDiskon, DapatCashback int
	
	fmt.Scan(&banyaknyaPembeli)
	DapatDiskon = 0
	DapatCashback = 0
	for i = 1; i <= banyaknyaPembeli; i++ {
		fmt.Scan(&totalBelanja)
		DapatDiskon = DapatDiskon + diskon(totalBelanja)
		DapatCashback = DapatCashback + cashback(totalBelanja)
	}
	fmt.Println(DapatDiskon,DapatCashback)
}

func diskon(totalBelanja int) int {
    if totalBelanja%100 == 0 || (totalBelanja%4 == 0 && totalBelanja%100 != 0) {
        return 1
    } else {
        return 0
    }
}

func cashback(totalBelanja int) int {
    if totalBelanja%100 == 0 || (totalBelanja%3 == 0 && totalBelanja%100 != 0) {
        return 1
    } else {
        return 0
    }
}

