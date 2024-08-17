package main
import "fmt"


func main() {
	var n, i, j int
	
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if i == j {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
