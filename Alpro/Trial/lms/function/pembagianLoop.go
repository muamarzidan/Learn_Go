package main
import "fmt"

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(bagi(bil1, bil2))
}

func bagi(bil1, bil2 int) int {
	var hasil int
	for bil1 >= bil2 { // bil 1 = 7, bil2 = 2
		bil1 -= bil2 // 7 - 2 = 5, 5 - 2 = 3, 3 - 2 = 1
		hasil++ // 1, 2, 3
	}
	return hasil
}

