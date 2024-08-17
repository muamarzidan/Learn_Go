package main
import "fmt"


type tabGoal struct {
	data [100]int
}

func inputData(t *tabGoal, n *int) {
	var x int
	*n = 0
	fmt.Scan(&x)
	for x >= 0 {
		t.data[*n] = x
		*n++
		fmt.Scan(&x)
	}
}

func rata(t tabGoal, n int) float64 {
	var sum int
	for i := 0; i < n; i++ {
		sum += t.data[i]
	}
	return float64(sum) / float64(n)
}

func main() {
	var tim1, tim2 tabGoal
	var n1, n2 int
	inputData(&tim1, &n1)
	inputData(&tim2, &n2)
	fmt.Println(rata(tim1, n1), rata(tim2, n2))
}