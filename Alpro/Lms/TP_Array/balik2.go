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
    /*
        IS: Array karakter (k) dan banyak elemen (n) terdefinisi sembarang
        FS: Array karakter (k) dan banyak elemen (n) terdefinisi
    */
    fmt.Scan(n)

    if *n > NMAX {
        *n = NMAX
    }

    for i := 0; i < *n; i++ {
        fmt.Scanf("%c", &k[i])
        fmt.Printf("%d %c\n", i, k[i])
    }
}

func cetak(k tabChar, n int) {
    /*
        IS: Array k dan n terdefinisi
        FS: Array k tercetak di layar dengan urutan terbalik
    */
    for i := n - 1; i >= 0; i-- {
        fmt.Printf("%c", k[i])
    }
}