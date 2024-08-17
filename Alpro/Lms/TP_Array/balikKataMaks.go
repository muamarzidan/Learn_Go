package main
import "fmt"


const NMAX int = 20
type tabChar [NMAX]byte

func main() {
    var kata tabChar
    var nChar int

    fmt.Scan(&nChar)
    baca(&kata, &nChar)
    cetak(kata, nChar)
}

func baca(k *tabChar, n *int) {
    var input string
    
    fmt.Scan(&input)
    if *n > NMAX {
        *n = NMAX
    }
    for i := 0; i < *n; i++ {
        k[i] = input[i]
		fmt.Printf("%d %c\n", i, k[i])
    }
}

func cetak(k tabChar, n int) {
    for i := n - 1; i >= 0; i-- {
        fmt.Printf("%c", k[i])
    }
}