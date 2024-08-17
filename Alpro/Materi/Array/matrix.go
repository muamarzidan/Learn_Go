package main 
import "fmt"	


const MAXINPUT int = 10
type maxMatrix [MAXINPUT]int

func main() {
	var input maxMatrix
	var n int
	fmt.Scan(&n)
	baca(&input, n)
	cetak(input, n)
}

func baca(input *maxMatrix, n int) {
	if n > MAXINPUT {
		n = MAXINPUT
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
}

func cetak(input maxMatrix, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", input[j])
		}
		fmt.Println()
	}
}