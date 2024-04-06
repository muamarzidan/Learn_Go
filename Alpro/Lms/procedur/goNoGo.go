package main

import "fmt"

func main() {
	var hujanTurun, bawaPayung string
	fmt.Scanln(&hujanTurun, &bawaPayung)
	goNoGo(hujanTurun, bawaPayung)
}

func goNoGo(hujanTurun, bawaPayung string) {
	if hujanTurun == "ya" {
		if bawaPayung == "ya" {
			fmt.Println("berangkat")
		} else {
			fmt.Println("diam di rumah")
		}
	} else {
		fmt.Println("berangkat")
	}
}