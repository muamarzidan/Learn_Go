// package main

// import "fmt"

// func main() {
// 	var bilangan int
// 	fmt.Scan(&bilangan)
// 	hitungNCari(bilangan)
// }

// func hitungNCari(bilangan int) {
// 	var i, nilaiTengah, total, digitJumlah int
// 	var pAwal, pTengah, pAkhir int
// 	digitJumlah = bilangan
// 	pAkhir = bilangan % 10

// 	for digitJumlah > 0 {
// 		total++
// 		digitJumlah = digitJumlah / 10
// 	}
// 	fmt.Println(total) 

// 	if total %2 == 0 {
// 		nilaiTengah = total / 2
// 	} else {
// 		nilaiTengah = total / 2 + 1 
// 	}

// 	i = 1
	
// 	for i <= total {
// 		if i == nilaiTengah {
// 			pTengah = bilangan % 10
// 		} else if i == total {
// 			pAwal = bilangan % 10
// 		}
// 		bilangan = bilangan / 10
// 		i++
// 	}

// 	fmt.Println(pAwal, pTengah, pAkhir)
// }

// package main

// import "fmt"

// // Definisikan struct untuk student
// type Student struct {
// 	Name  string
// 	Age   int
// 	Grade int
// }

// func main() {
// 	// Deklarasi tiga buah data mahasiswa (students) menggunakan struct Student
// 	var student1, student2, student3 Student
// 	student1 = Student{
// 		Name:  "Ahmad",
// 		Age:   20,
// 		Grade: 85}
// 	student2 = Student{
// 		Name:  "Doni",
// 		Age:   21,
// 		Grade: 90}
// 	student3 = Student{
// 		Name:  "Putri",
// 		Age:   19,
// 		Grade: 75}

// 	// Menampilkan data mahasiswa
// 	fmt.Println("Data Mahasiswa 1:")
// 	fmt.Println("Nama:", student1.Name)
// 	fmt.Println("Umur:", student1.Age)
// 	fmt.Println("Grade:", student1.Grade)

// 	fmt.Println("\nData Mahasiswa 2:")
// 	fmt.Println("Nama:", student2.Name)
// 	fmt.Println("Umur:", student2.Age)
// 	fmt.Println("Grade:", student2.Grade)

// 	fmt.Println("\nData Mahasiswa 3:")
// 	fmt.Println("Nama:", student3.Name)
// 	fmt.Println("Umur:", student3.Age)
// 	fmt.Println("Grade:", student3.Grade)
// }


// PSEUDOCODE

// struct student <
// 	Name : string
// 	Age : int
// 	Grade : int >

// program main

// kamus
// 	student1, student2, student3 : student

// algoritma
// 	output("Data Mahasiswa 1:")
// 	output("Nama:", student1.Name)
// 	output("Umur:", student1.Age)
// 	output("Grade:", student1.Grade)

// 	output("Data Mahasiswa 2:")
// 	output("Nama:", student1.Name)
// 	output("Umur:", student1.Age)
// 	output("Grade:", student1.Grade)

// 	output("Data Mahasiswa 1:")
// 	output("Nama:", student1.Name)
// 	output("Umur:", student1.Age)
// 	output("Grade:", student1.Grade)



// func cetakPola() {
// 	var N int
// 	fmt.Scanln(&N)
// 	pola(1, N)
// }

// func pola(baris, n int) int {
// 	if baris <= n {
// 		cetakBintang(baris)
// 		fmt.Println()
// 		return pola(baris+1, n)
// 	}
// 	return 0

// }

// func cetakBintang(jumlah int) {
// 	if jumlah > 0 {
// 		fmt.Print("* ")
// 		cetakBintang(jumlah - 1)
// 	}
// }



package main

import "fmt"

// Definisikan struct untuk student
type Student struct {
	Name  string
	Age   int
	Grade [3]int
}

func main() {
	var student1 Student

	// Mengisi data mahasiswa
	fmt.Println("Masukkan nama mahasiswa:")
	fmt.Scanln(&student1.Name)
	fmt.Println("Masukkan usia mahasiswa:")
	fmt.Scanln(&student1.Age)
	fmt.Println("Masukkan nilai-nilai Grade (sebanyak 6):")
	for i := 0; i < 2; i++ {
		fmt.Printf("Nilai Grade %d: ", i+1)
		fmt.Scanln(&student1.Grade[i])
	}
	// Menampilkan data mahasiswa
	fmt.Println("Tabel Data Mahasiswa:")
	fmt.Println("Nama:", student1.Name)
	fmt.Println("Umur:", student1.Age)
	for i := 0; i < 2; i++ {
		fmt.Printf("Grade %d: %d\n", i+1, student1.Grade[i])
	}
}

// PSEUDOCODE

// struct student <
// 	Name : string
// 	Age : int
// 	Grade : array[3] of integer >

// program main

// kamus
// 	student1 : student

// algoritma
// 	output("Masukkan nama mahasiswa:")
// 	input(student1.Name)
// 	output("Masukkan usia mahasiswa:")
// 	input(student1.Age)
// 	output("Masukkan nilai-nilai Grade (sebanyak 6):")
// 	for i <- 0 to 2 do
// 		output("Nilai Grade ", i+1, ": ")
// 		input(student1.Grade[i])
// 	endfor

// 	output("Tabel Data Mahasiswa:")
// 	output("Nama:", student1.Name)
// 	output("Umur:", student1.Age)
// 	for i <- 0 to 2 do
// 		output("Grade ", i+1, ": ", student1.Grade[i])
// 	endfor
// endprogram