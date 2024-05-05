package main 
import "fmt"

const POKEMONMAX = 1000
const NMAX = 3


type pokemon struct  {
	nama string
	namaPokemon string 
	jumlahPokemon int
	kelangkaan string 
	elemen string
}

type tasPokemon [POKEMONMAX]pokemon

func main() {
	var data tasPokemon
	fmt.Println("==================================================")
	fmt.Println("= Tulis Tingkat Kelangkaan : rare, medium, low   =")
	fmt.Println("= Tulis Elemen : Listrik, Tanah, Api, Air, Angin =")
	fmt.Println("==================================================")
	fmt.Print("\n")

	inputData(&data)

	fmt.Println("Selamat kepada :", cariNamaOrangPalingBanyakPokemon(data), "telah mendapatkan pokemon paling banyak dengan jumlah pokemon sebanyak :", data[cariPalingBanyakPokemon(data)].jumlahPokemon)
	fmt.Println("------Dengan Spesifikasi Pokemon------")
	fmt.Println("Nama Pokemon :", data[cariPalingBanyakPokemon(data)].namaPokemon)
	fmt.Println("Tingkat Kelangkaan :", data[cariPalingBanyakPokemon(data)].kelangkaan)
	fmt.Println("Elemen Pokemon :", data[cariPalingBanyakPokemon(data)].elemen)
	fmt.Println("---------------------------------------")

	fmt.Println("Namun sayang sekali, kepada :", cariNamaOrangPalingSedikitPokemon(data), "telah mendapatkan pokemon paling sedikit dengan jumlah pokemon sebanyak :", data[cariPalingSedikitPokemon(data)].jumlahPokemon)
}

func inputData(data *tasPokemon) {
	for i := 0; i < NMAX; i++ {
		fmt.Println("Silahkan masukan data orang ke", i+1)
		fmt.Println("Masukan nama anda : ")
		fmt.Scanln(&data[i].nama)
		fmt.Println("Masukan nama pokemon yang diperoleh : ")
		fmt.Scanln(&data[i].namaPokemon)
		fmt.Println("Masukan jumlah pokemon yang diperoleh : ")
		fmt.Scanln(&data[i].jumlahPokemon)
		fmt.Println("Masukan tingkat kelangkaan pokemon : ")
		fmt.Scanln(&data[i].kelangkaan)
		fmt.Println("Masukan elemen pokemon : ")
		fmt.Scanln(&data[i].elemen)
		fmt.Print("\n")

		if data[i].jumlahPokemon > POKEMONMAX {
			data[i].jumlahPokemon = POKEMONMAX
		}
	}
}

func cariPalingBanyakPokemon (data tasPokemon) int {
    var max int
    var idx int
    for i := 0; i < NMAX; i++ {
        if data[i].jumlahPokemon > max {
            max = data[i].jumlahPokemon
            idx = i
        }
    }
    return idx
}

func cariNamaOrangPalingBanyakPokemon (data tasPokemon) string {
	var max int
	var nama string
	for i := 0; i < NMAX; i++ {
		if data[i].jumlahPokemon > max {
			max = data[i].jumlahPokemon
			nama = data[i].nama
		}
	}
	return nama
}

func cariPalingSedikitPokemon (data tasPokemon) int {
    var min int = data[0].jumlahPokemon
    var idx int
    for i := 0; i < NMAX; i++ {
        if data[i].jumlahPokemon < min {
            min = data[i].jumlahPokemon
            idx = i
        }
    }
    return idx
}

func cariNamaOrangPalingSedikitPokemon (data tasPokemon) string {
	var min int
	var nama string
	for i := 0; i < NMAX; i++ {
		min = data[i].jumlahPokemon
		nama = data[i].nama
		if data[i].jumlahPokemon < min {
			min = data[i].jumlahPokemon
			nama = data[i].nama
		}
	}
	return nama
}



// PSEUDOCODE 

// constant POKEMONMAX : integer = 5

// type pokemon <
// 	nama : string
// 	namaPokemon : string
// 	jumlahPokemon : integer
// 	kelangkaan : string
// 	elemen : string
// >

// type tasPokemon : array [0...POKEMONMAX-1] of pokemon

// program main 

// kamus 
// 	data : tasPokemon
// 	index : integer

// algoritma
// 	for i <- 0 to POKEMONMAX-1 do
// 		output("Silahkan masukan data orang ke", i+1)
// 		output("Masukan nama anda : ")
// 		input(data[i].nama)
// 		output("Masukan nama pokemon yang diperoleh : ")
// 		input(data[i].namaPokemon)
// 		output("Masukan jumlah pokemon yang diperoleh : ")
// 		input(data[i].jumlahPokemon)
// 		output("Masukan tingkat kelangkaan pokemon : ")
// 		input(data[i].kelangkaan)
// 		output("Masukan elemen pokemon : ")
// 		input(data[i].elemen)
// 		output("\n")
// 	endfor

// 	index <- cariPalingBanyakPokemon(data)
// 	output("Selamat kepada :", data[index].nama, "telah mendapatkan pokemon paling banyak dengan jumlah pokemon sebanyak :", data[index].jumlahPokemon)
// 	output("------Dengan Spesifikasi Pokemon------")
// 	output("Nama Pokemon :", data[index].namaPokemon)
// 	output("Tingkat Kelangkaan :", data[index].kelangkaan)
// 	output("Elemen Pokemon :", data[index].elemen)
// 	output("---------------------------------------")

// 	index <- cariPalingSedikitPokemon(data)
// 	output("Namun sayang sekali, kepada :", data[index].nama, "telah mendapatkan pokemon paling sedikit dengan jumlah pokemon sebanyak :", data[index].jumlahPokemon)

// endprogram

// function cariPalingBanyakPokemon (data : tasPokemon) -> integer
// { pada fungsi ini akan mencari data orang yang mendapatkan pokemon paling banyak }

// kamus 
// 	max, index, i : integer

// algoritma
// 	max <- 0
// 	index <- 0
// 	for i <- 0 to POKEMONMAX-1 do
// 		if data[i].jumlahPokemon > max then
// 			max <- data[i].jumlahPokemon
// 			index <- i
// 		endif
// 	endfor
// 	return index
// endfunction


// function cariPalingSedikitPokemon (data : tasPokemon) -> integer
// { pada fungsi ini akan mencari data orang yang mendapatkan pokemon paling sedikit }

// kamus
// 	min, index, i : integer

// algoritma
// 	min <- data[0].jumlahPokemon
// 	index <- 0
// 	for i <- 0 to POKEMONMAX-1 do
// 		if data[i].jumlahPokemon < min then
// 			min <- data[i].jumlahPokemon
// 			index <- i
// 		endif
// 	endfor
// 	return index

// endfunction



















	// cariPalingBanyakPokemon(data)
	// fmt.Print("\n")
	// cariPalingSedikitPokemon(data)

		// var max int
	// var nama string
	// var namaPokemon string
	// var kelangkaan string
	// var elemen string
	// fmt.Println("Selamat kepada :", nama, "telah mendapatkan pokemon paling banyak dengan jumlah pokemon sebanyak :", max)
	// fmt.Println("------Dengan Spesifikasi Pokemon------")
	// fmt.Println("Nama Pokemon :", namaPokemon)
	// fmt.Println("Tingkat Kelangkaan :", kelangkaan)
	// fmt.Println("Elemen Pokemon :", elemen)
	// fmt.Println("---------------------------------------")
	// for i := 0; i < POKEMONMAX; i++ {
	// 	if data[i].jumlahPokemon > max {
	// 		max = data[i].jumlahPokemon
	// 		nama = data[i].nama
	// 		namaPokemon = data[i].namaPokemon
	// 		kelangkaan = data[i].kelangkaan
	// 		elemen = data[i].elemen
	// 	}
	// }

	// var min int
	// var nama string 
	// for i := 0; i < POKEMONMAX; i++ {
	// 	min = data[i].jumlahPokemon
	// 	nama = data[i].nama
	// 	if data[i].jumlahPokemon < min {
	// 		min = data[i].jumlahPokemon
	// 		nama = data[i].nama
	// 	}
	// }
	// fmt.Println("Namun sayang sekali, kepada :", nama, "telah mendapatkan pokemon paling sedikit dengan jumlah pokemon sebanyak :", min)

