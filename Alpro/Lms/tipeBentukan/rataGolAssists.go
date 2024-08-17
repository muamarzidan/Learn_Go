package main
import "fmt"


func main() {
	type pemain struct {
		nama   string
		gol    int
		assist int
	}

	var a, b, c pemain

	fmt.Scan(&a.nama, &a.gol, &a.assist)
	fmt.Scan(&b.nama, &b.gol, &b.assist)
	fmt.Scan(&c.nama, &c.gol, &c.assist)

	var rataGol, rataAssist float64
	rataGol = float64(a.gol+b.gol+c.gol) / 3
	rataAssist = float64(a.assist+b.assist+c.assist) / 3
	fmt.Println(rataGol)
	fmt.Println(rataAssist)
}
