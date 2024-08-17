package main 
import "fmt"

const nMax int = 1000
type tabChar [nMax]rune

func isiArray(arr *tabChar, n *int) {
	var kata string
	var i int
	i = 0

	fmt.Scan(&kata)
	for i < len(kata) {
		(*arr)[i] = rune(kata[i])
		*n = *n + 1
		i = i + 1
	}
}

func isSudahAda(char rune, arr tabChar, total int) bool {
	var i int
	for i = 0; i < total; i++ {
		if arr[i] == char {
			return true
		}
	}
	return false
}

func isVokal(char rune) bool {
	return char == 'A' || char == 'a' || char == 'I' || char == 'i' || char == 'U' || char == 'u' || char == 'E' || char == 'e' || char == 'O' || char == 'o'
}

func prosesHurufVokal(arr tabChar, n int, total *int, arrVokal *tabChar) {
	*total = 0
    var i int
	for i = 0; i < n; i++ {
		if !isSudahAda(arr[i], *arrVokal, *total) && isVokal(arr[i]) {
			(*arrVokal)[*total] = arr[i]
			*total = *total + 1
		}
	}
}

func ProsesHurufKonsonan(arr tabChar, n int, total *int, arrKonsonan *tabChar) {
	*total = 0
	var i int
	for i = 0; i < n; i++ {
		if !isSudahAda(arr[i], *arrKonsonan, *total) && !isVokal(arr[i]) && ((arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= 'a' && arr[i] <= 'z')) {
			(*arrKonsonan)[*total] = arr[i]
			*total = *total + 1
		}
	}
}

func ProsesHurufKarakter(arr tabChar, n int, total *int, arrKar *tabChar) {
	*total = 0
	var i int
	for i = 0; i < n; i++ {
		if !isSudahAda(arr[i], *arrKar, *total) && (arr[i] < 'A' || (arr[i] > 'Z' && arr[i] < 'a') || arr[i] > 'z') {
			(*arrKar)[*total] = arr[i]
			*total = *total + 1
		}
	}
}

func cetakHuruf(arr tabChar, n int) {
	var totalVokal, totalKonsonan, totalKarakterLain, i int
	var arrVokal, arrKonsonan, arrKarakterLain tabChar

	prosesHurufVokal(arr, n, &totalVokal, &arrVokal)
	ProsesHurufKonsonan(arr, n, &totalKonsonan, &arrKonsonan)
	ProsesHurufKarakter(arr, n, &totalKarakterLain, &arrKarakterLain)

	fmt.Println(totalVokal)
	if totalVokal == 0 {
		fmt.Print("Tidak ditemukan")
	} else {
		for i = 0; i < totalVokal; i++ {
			fmt.Printf("%c ", arrVokal[i])
		}
	}

	fmt.Println()
	fmt.Print(totalKonsonan)
	fmt.Println()
	if totalKonsonan == 0 {
		fmt.Print("Tidak ditemukan")
	} else {
		for i = 0; i < totalKonsonan; i++ {
			fmt.Printf("%c ", arrKonsonan[i])
		}
	}

	fmt.Println()
	fmt.Print(totalKarakterLain)
	fmt.Println()
	if totalKarakterLain == 0 {
		fmt.Print("Tidak ditemukan")
	} else {
		for i = 0; i < totalKarakterLain; i++ {
			fmt.Printf("%c ", arrKarakterLain[i])
		}
	}
}

func main() {
	var n int
	var arr tabChar
	
	isiArray(&arr, &n)
	cetakHuruf(arr, n)
}