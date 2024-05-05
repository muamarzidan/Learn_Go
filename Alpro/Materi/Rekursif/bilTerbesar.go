package main  

import "fmt"

const END_INPUT = -1

func main() {
	var inputVal int
	fmt.Scan(&inputVal)
	if inputVal != END_INPUT {
		findMaxValue(inputVal)
	} else {
		fmt.Println("Program selesai", inputVal)
	}
}

func findMaxValue(inputVal int) int {
	var x int
	fmt.Scan(&x)
	if x == END_INPUT {
		return inputVal
	} else if x > inputVal {
		return findMaxValue(x)
	} else {
		return findMaxValue(inputVal)
	}
}