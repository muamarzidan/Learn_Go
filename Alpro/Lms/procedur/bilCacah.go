package main
import "fmt"


func main() {
	var bil int

	fmt.Scanln(&bil)
	cacah(bil)
}

func cacah(bilangan int) {
	for bilangan != 0 {
		fmt.Printf("%d ", bilangan%10)
		bilangan /= 10
	}
}