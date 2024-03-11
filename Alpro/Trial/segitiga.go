package main
import "fmt"

// segitiga siku-siku
func main() {
	var input int
	fmt.Scan(&input)
	for i:= 0; i <= input; i ++ {
		for j := 0; j < i; j++ {
				fmt.Print("X")
		}
		fmt.Println()
	}
}

// segitiga siku-siku terbalik
func main() {
	var input int
	fmt.Scan(&input)
	for i:= 0; i <= input; i ++ {
		for j := 0; j < input - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
				fmt.Print("X")
		}
		fmt.Println()
	}
}

// segitiga sama sisi
func main() {
	var input int
	fmt.Scan(&input)
	for i:= 0; i < input; i ++ {
		for j := 0; j < input - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
				fmt.Print("X")
		}
		for j := 0; j < i - 1; j++ {
				fmt.Print("X")
		}
		fmt.Println()
	}
}