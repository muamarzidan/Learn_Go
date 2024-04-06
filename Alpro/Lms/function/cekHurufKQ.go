package main

import "fmt"

func main() {
	var hrf byte
	var ada bool
	fmt.Scanf("%c", &hrf)

	for hrf != '.' && !ada {
		ada = huruf(hrf)
		fmt.Scanf("%c", &hrf)
	}
	
	fmt.Println(ada)
}

func huruf(hrf byte) bool {
	return hrf == 'k' || hrf == 'q'
}


