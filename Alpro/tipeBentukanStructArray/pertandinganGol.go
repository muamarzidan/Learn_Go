// Sebuah program digunakan untuk menyimpan jumlah kemenangan pertandingan bola 2 tim dalam setiap tahunnya.

// Masukan terdiri dari 2 baris. Baris pertama adalah kumpulan bilangan bulat yang menyatakan gol-gol yang diperoleh tim 1, 
// sedangkan baris kedua adalah kumpulan gol yang diperoleh tim 2. Masukan pada setiap barisnya berakhir apabila bilangan 
// yang diberikan adalah negatif. 

// Keluaran berupa 2 bilangan yang menyatakan rata-rata kemenangan dari 2 tim tersebut setiap tahunnya.


// ada 2 fungsi selain fungi main yaitu :

// procedure inputData(in/out t: tabGoal, n : integer)
// { I.S. : data kemenangan suatu tim telah siap pada piranti masukan}
// { F.S. : t berisi n data kemanaan tim}

// function rata(t: tabGoal, n: integer) -> real
// {diberikan array t yang berisi n data kemenangan tim, fungsi ini mengembalikan rata-rata kemenangan tim}

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