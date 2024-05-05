package main 
import "fmt"

// func main() {
// 	var jumlah int
//     jumlah = 0
//     kelipatanTiga(&jumlah)
//     fmt.Println(jumlah)
// }

// func kelipatanTiga(jumlah *int){
//     /*I.S. jumlah diinisialisasi dengan nilai 0
//     F.S. jumlah telah berisi jumlah dari bilangan-bilangan ganjil kelipatan 3
//     */
//     var n int
//     fmt.Scan(&n)
    
//     for n >= 0 {
//         if n % 3 == 0{
//             *jumlah = *jumlah + n
// 			fmt.Println(*jumlah)
//         }
// 		fmt.Scan(&n)
//     }
// }


// Buatlah procedure dalam bahasa Go bernama cariGo untuk menghitung banyaknya string "go" pada serangkaian karakter yang diinputkan. 

// procedure cariGo(in kar : char){

//    IS: terdefinisi kar berisi serangkaian karakter

//    FS: program mengeluarkan jumlah karakter go yang ada dalam karakter

// }

// Masukan terdiri dari serangkaian karakter dengan diakhiri tanda titik. 

// Keluaran berupa bilangan bulat yang menyatakan banyaknya string "go".

// Contoh Masukan
// inigolang.
// Contoh Keluaran
// 1

// Contoh Masukan
// inigolanginggolang.
// Contoh Keluaran
// 2

// Contoh Masukan
// golangolang.
// Contoh Keluaran
// 2


// func main() {
// 	var kar byte
// 	fmt.Scanf("%c", &kar)
// 	cariGo(kar)
// }

// func cariGo(kar byte) {
//     var tot int
//     tot = 0
//     var awal byte
//     for kar != '.' {
//         if awal == 'g' && kar == 'o' {
//             tot++
//         }
//         awal = kar
//         fmt.Scanf("%c", &kar)
//     }
//     fmt.Println(tot)
// }



// Hitunglah berapa banyaknya huruf kecil atau lowercase dari sebuah kata.

// Masukan berupa string kata 

// Keluaran jumlah huruf lowercase dari kata tersebut

// contoh masukan
// Halo
// contoh keluaran
// 3

// contoh masukan
// HaloDunia
// contoh keluaran
// 7
// contoh masukan
// saya 
// contoh keluaran
// 4

// contoh masukan 
// s13a
// contoh keluaran
// 2

// func main() {
//     var kata string
// 	var jumlah, index int
// 	index = 0
// 	fmt.Scanln(&kata)
// 	jumlah = hitungLowercase(kata, index)
// 	fmt.Print(jumlah)
// }

// func hitungLowercase(kata string, index int) int {
//     /*{ I.S. Terdefinisi parameter kata sebagai input string
//     F.S. menampilkan jummlah  jumlah huruf lowercase dari kata menggunakan fungsi rekursif*/
//     if index == len(kata) { // masukan kondisi untuk mengecek index dengan jumlah string
//         return 0
//     }
//     if 'a' <= kata[index] && kata[index] <= 'z' {
//         return 1 + hitungLowercase(kata, index+1) // gunakan fungsi rekursif yang tepat
//     }
//     return hitungLowercase(kata, index+1)
// }



Buatlah sebuah fungsi rekursif untuk mengembalikan sebuah bilangan bulat yang menyatakan hasil perkalian bilangan ganjil dari 1 hingga N.

Masukan berupa sebuah bilangan bulat positif N.
Keluaran berupa sebuah bilangan bulat yang menyatakan hasil perkalian bilangan ganjil dari 1 hingga N.

Contoh Masukan
1
Contoh Keluaran
1

Contoh Masukan
3
Contoh Keluaran
3

Contoh Masukan
5
Contoh Keluaran
15


func main() {
    var N int
    fmt.Scan(&N)
    fmt.Println(perkalianBilanganGanjil(N))
}

func perkalianBilanganGanjil(... int) int {
    /*  fungsi mengembalikan sebuah bilangan bulat yang menyatakan 
        hasil perkalian bilangan ganjil dari 1 hingga N */
        
        if x == 1 {
            return 1
        } else if x%2 == 0 {
            return ___ // masukkan fungsi rekursif pada baris ini 
        }
        return ___ // masukkan fungsi rekursif pada baris ini 
    }

lengkap code diaats 

func perkalianBilanganGanjil(x int) int {
    /*  fungsi mengembalikan sebuah bilangan bulat yang menyatakan 
        hasil perkalian bilangan ganjil dari 1 hingga N */
        
        if x == 1 {
            return 1
        } else if x%2 == 0 {
            return perkalianBilanganGanjil(x-1)
        }
        return x * perkalianBilanganGanjil(x-1)
    }

