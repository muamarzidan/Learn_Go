package main
import "fmt"


func main() {
	var pilih int
	
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)?")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		case 4:
			return
		}
	}
}

func menu() {
	fmt.Println("----------------------")
	fmt.Println("\tM E N U")
	fmt.Println("----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("----------------------")
}

func hitungJumlah() {
	var a, b int
	fmt.Print("Masukan dua bilangan yang akan dijumlahkan : ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil penjumlahan : ", a+b)
}

func hitungKali() {
	var a, b int
	fmt.Print("Masukan dua bilangan yang akan dikalikan : ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian : ", a*b)
}

func hitungBagi() {
	var a, b float64
	fmt.Print("Masukan dua bilangan yang akan dibagi : ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil pembagian : ", a/b)
}
