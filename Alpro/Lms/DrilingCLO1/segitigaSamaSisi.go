package main 

import "fmt"

func main() {
	var s1, s2, s3 float64
	var segitigaSS bool
	fmt.Scan(&s1, &s2, &s3)
	segitigaSS = segitigaSamaSisi(s1, s2, s3)
	fmt.Println(segitigaSS)
}

// Buat fungsi saja
func segitigaSamaSisi(s1, s2, s3 float64) bool {
    // Mengembalikan nilai boolean jika 3 bilangan real dapat
    // membentuk segitiga sama sisi
    return (s1 == s2 && s2 == s3) && s1 != 0 && s2 != 0 && s3 != 0
}