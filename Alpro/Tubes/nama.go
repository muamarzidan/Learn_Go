package main

import (
    "fmt"
    "strings"
)

func main() {
    var nama []string
    fmt.Print("Masukkan Nama: ")
    fmt.Scan(&nama)

    namaLengkap := strings.Join(nama, " ")
    fmt.Println("Nama yang dimasukkan adalah:", namaLengkap)
}