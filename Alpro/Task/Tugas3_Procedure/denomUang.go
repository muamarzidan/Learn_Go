package main

import "fmt"

func main() {
	var uang, hasil, sisa int
	fmt.Scanf("%d", &uang)

	denomUang(&hasil, &sisa, uang)
}

func denomUang(hasil *int, sisa *int, uang int) {
	const (
		sepuluhRibu = 10000
		duaRibu     = 2000
		seribu      = 1000
		lembar      = " lembar"
	)

	*hasil = uang / sepuluhRibu
	*sisa = uang % sepuluhRibu
	fmt.Println(*hasil, lembar, "10000")

	*hasil = *sisa / duaRibu
	*sisa = *sisa % duaRibu
	fmt.Println(*hasil, lembar, "2000")

	*hasil = *sisa / seribu
	*sisa = *sisa % seribu
	fmt.Println(*hasil, lembar, "1000")
}


// PSEUDOCODE
// program main

// kamus
// 	uang, hasil, sisa : integer

// algoritma
// 	input(uang)
// 	denomUang(hasil, sisa, uang)

// enedprogram

// procedure denomUang(in uang : integer, in/out hasil : integer, in/out sisa : integer)

// {IS : initial state terdefinisi adalah uang dengan bilangan bulat positif}
// {FS : final state terdefinisi adalah hasil dan sisa dengan bilangan bulat positif}

// kamus
// 	constant sepuluhRibu : integer = 10000
// 	constant duaRibu : integer = 2000
// 	constant seribu : integer = 1000
// 	constant lembar : string = " lembar"

// algoritma
// 	hasil <- uang div sepuluhRibu
// 	sisa <- uang mod sepuluhRibu

// 	output(hasil, lembar, "10000")

// 	hasil <- sisa div duaRibu
// 	sisa <- sisa mod duaRibu

// 	output(hasil, lembar, "2000")

// 	hasil <- sisa div seribu
// 	sisa <- sisa mod seribu

// 	output(hasil, lembar, "1000")
// endprocedure
