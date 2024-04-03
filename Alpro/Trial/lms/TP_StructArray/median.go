package main

import "fmt"

type array [100]int

func main() {
	var arr array
	var n int
	isiArray(&n, &arr)
	fmt.Print("Median : ", median(n, arr))
}

func isiArray(n *int, arr *array) {
	/*I.S. Terdefinisi pointer n bilangan bulat dan pointer array untuk menampung bilangan bulat
	F.S. n terisi jumlah data yang diinput dan pointer array terisi bilangan bulat sebanyak n*/
	fmt.Scan(n)
	// fmt.Print("Masukkan data ke- ", n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&arr[i])
	}
}

func median(n int, arr array) float64 {
	/* mengembalikan median atau nilai tengah dari array */
	var median float64
	if n % 2 == 0 {
		n = n / 2
		median = (float64(arr[n-1]) + float64(arr[n])) / 2
	} else {
		n = (n + 1) / 2
		median = (float64(arr[n-1]) + float64(arr[n])) / 2
	}
	return median
}
