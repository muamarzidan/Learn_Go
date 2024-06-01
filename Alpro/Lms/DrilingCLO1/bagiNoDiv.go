package main

import "fmt"

func main() {
    var M, N int
    fmt.Scan(&M, &N)
    fmt.Println(pembagianInteger(M, N))
}

func pembagianInteger(M int, N int) int {
    if M < N {
        return 0
    }
    return 1 + pembagianInteger(M-N, N)
}