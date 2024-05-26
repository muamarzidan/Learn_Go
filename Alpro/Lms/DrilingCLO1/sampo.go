package main 

import "fmt"

func main() {
	var habisSampo, hujan string
	fmt.Scan(&habisSampo)
	fmt.Scan(&hujan)
	haveSampoHujan(habisSampo, hujan)
}

func haveSampoHujan(habisSampo string, hujan string) {
    //Menampilkan pergi atau tidak pergi ke minimarket sesuai dengan ketentuan soal
	if habisSampo == "ya" && hujan == "ya" {
		fmt.Println("tidak pergi ke minimarket")
	} else if habisSampo == "ya" && hujan == "tidak" {
		fmt.Println("pergi ke minimarket")
	} else if habisSampo == "tidak" {
		fmt.Println("tidak pergi ke minimarket")
	}
}