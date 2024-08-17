package main 
import "fmt"


const POKEMONMAX = 1000
const NMAX = 5

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

// constant NMAX : integer = 5
// constant POKEMONMAX : integer = 1000

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
// 	output("==================================================")
// 	output("= Tulis Tingkat Kelangkaan : rare, medium, low   =")
// 	output("= Tulis Elemen : Listrik, Tanah, Api, Air, Angin =")
// 	output("==================================================")
// 	output("\n")

// 	inputData(data)

// 	output("Selamat kepada :", cariNamaOrangPalingBanyakPokemon(data), "telah mendapatkan pokemon paling banyak dengan jumlah pokemon sebanyak :", data[cariPalingBanyakPokemon(data)].jumlahPokemon)
// 	output("------Dengan Spesifikasi Pokemon------")
// 	output("Nama Pokemon :", data[cariPalingBanyakPokemon(data)].namaPokemon)
// 	output("Tingkat Kelangkaan :", data[cariPalingBanyakPokemon(data)].kelangkaan)
// 	output("Elemen Pokemon :", data[cariPalingBanyakPokemon(data)].elemen)
// 	output("---------------------------------------")

// 	output("Namun sayang sekali, kepada :", cariNamaOrangPalingSedikitPokemon(data), "telah mendapatkan pokemon paling sedikit dengan jumlah pokemon sebanyak :", data[cariPalingSedikitPokemon(data)].jumlahPokemon)
// endprogram


// procedure inputData (data : tasPokemon)
// {
// 	pada procedure ini, user akan diminta untuk memasukkan data sebanyak NMAX kali, NMAX ini sebagai batas maksimal input data
// }

// kamus
// 	i : integer

// algoritma	
// 	for i <- 0 to NMAX-1 do
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

// 		if data[i].jumlahPokemon > POKEMONMAX then
// 			data[i].jumlahPokemon <- POKEMONMAX
// 		endif
// 	endfor

// endprocedure


// function cariPalingBanyakPokemon (data : tasPokemon) -> integer
// {
// 	pada fungsi ini, akan dicari data orang yang mendapatkan pokemon paling banyak
// }

// kamus
// 	max : integer
// 	idx : integer
// 	i : integer

// algoritma
// 	max <- data[0].jumlahPokemon
// 	idx <- 0
// 	for i <- 0 to NMAX-1 do
// 		if data[i].jumlahPokemon > max then
// 			max <- data[i].jumlahPokemon
// 			idx <- i
// 		endif
// 	endfor
// 	return idx

// endfunction


// function cariNamaOrangPalingBanyakPokemon (data : tasPokemon) -> string
// {
// 	pada fungsi ini, akan dicari nama orang yang mendapatkan pokemon paling banyak 
// }

// kamus
// 	max : integer
// 	nama : string
// 	i : integer

// algoritma
// 	max <- data[0].jumlahPokemon
// 	nama <- data[0].nama
// 	for i <- 0 to NMAX-1 do
// 		if data[i].jumlahPokemon > max then
// 			max <- data[i].jumlahPokemon
// 			nama <- data[i].nama
// 		endif
// 	endfor
// 	return nama

// endfunction


// function cariPalingSedikitPokemon (data : tasPokemon) -> integer
// {
// 	pada fungsi ini, akan dicari data orang yang mendapatkan pokemon paling sedikit
// }

// kamus
// 	min : integer
// 	idx : integer
// 	i : integer

// algoritma
// 	min <- data[0].jumlahPokemon
// 	idx <- 0
// 	for i <- 0 to NMAX-1 do
// 		if data[i].jumlahPokemon < min then
// 			min <- data[i].jumlahPokemon
// 			idx <- i
// 		endif
// 	endfor
// 	return idx

// endfunction


// function cariNamaOrangPalingSedikitPokemon (data : tasPokemon) -> string
// {
// 	pada fungsi ini, akan dicari nama orang yang mendapatkan pokemon paling sedikit
// }

// kamus
// 	min : integer
// 	nama : string
// 	i : integer

// algoritma
// 	min <- data[0].jumlahPokemon
// 	nama <- data[0].nama
// 	for i <- 0 to NMAX-1 do
// 		min <- data[i].jumlahPokemon
// 		nama <- data[i].nama
// 		if data[i].jumlahPokemon < min then
// 			min <- data[i].jumlahPokemon
// 			nama <- data[i].nama
// 		endif
// 	endfor
// 	return nama

// endfunction