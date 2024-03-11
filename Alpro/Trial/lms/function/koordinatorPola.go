package main

import ("fmt" ; "math")

func ukurX(R, T float64) float64 {
	var hasil float64
	hasil = R * math.Cos((T*math.Pi)/180)
	return hasil
}

func ukurY(R, T float64) float64 {
	var hasil float64
	hasil = R * math.Sin((T*math.Pi)/180)
	return hasil
}

func main() {
	var R, T, hasilSatu, hasilDua float64
	fmt.Scanf("%f %f", &R, &T)
	hasilSatu = ukurX(R, T)
	hasilDua = ukurY(R, T)

	fmt.Printf("%.2f %.2f\n", hasilSatu, hasilDua)
}
