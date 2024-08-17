package main
import "fmt"


func main() {
    var N int
	
    fmt.Scan(&N)
    fmt.Println(jumlahDeretGenap(N))
}

func jumlahDeretGenap(x int) int {
    if x == 0 {
		return 0
	} else if x%2 != 0 {
		return jumlahDeretGenap(x-1)
	}

	return x + jumlahDeretGenap(x-1)
}