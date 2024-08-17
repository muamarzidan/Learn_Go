package main
import "fmt"


func main() {
	var a, r, n int
	fmt.Scan(&a, &r, &n)
	fmt.Println(UnGeometri(a, r, n))
}

func UnGeometri(a, r, n int) int {
	if n == 1 {
		return a
	} else {
		return r * UnGeometri(a, r, n-1)
	}
}