package main
import "fmt"


type buku struct {
	title string
	author string
	page int
	price float64
}

func main() {
	buku1 := buku{
		title: "Belajar Go",
		author: "Programmer Zaman Now",	
		page: 200,
		price: 100000,
	}

	buku2 := buku{
		title: "Belajar Python",
		author: "DR Angel Ayu",
		page: 400,
		price: 300000,
	}

	buku3 := buku{
		title: "Belajar Javascript",
		author: "Maximilian Schwarzmüller",
		page: 300,
		price: 200000,
	}

	fmt.Println("Data Buku 1")
	fmt.Println("Title : ", buku1.title)
	fmt.Println("Author : ", buku1.author)
	fmt.Println("Page : ", buku1.page)
	fmt.Println("Price : ", buku1.price)
	fmt.Print("\n")
	fmt.Println("Data Buku 2")
	fmt.Println("Title : ", buku2.title)
	fmt.Println("Author : ", buku2.author)
	fmt.Println("Page : ", buku2.page)
	fmt.Println("Price : ", buku2.price)
	fmt.Print("\n")
	fmt.Println("Data Buku 3")
	fmt.Println("Title : ", buku3.title)
	fmt.Println("Author : ", buku3.author)
	fmt.Println("Page : ", buku3.page)
	fmt.Println("Price : ", buku3.price)
}


// TUGAS BUKU WITH ARRAY

type bukuArray struct {
	title    string
	author   string
	page     int
	price    float64
	kategori [4]string
}

func main() {
	buku1 := bukuArray{
		title:    "Belajar Go",
		author:   "Programmer Zaman Now",
		page:     200,
		price:    100000,
		kategori: [4]string{"Basic", "Middle", "Advance", "Expert"},
	}

	fmt.Println("Data Buku :")
	fmt.Println("Title : ", buku1.title)
	fmt.Println("Author : ", buku1.author)
	fmt.Println("Page : ", buku1.page)
	fmt.Println("Price : ", buku1.price)
	for i := 0; i < len(buku1.kategori); i++ {
		fmt.Println("Kategori : ", buku1.kategori[i])
	}
}