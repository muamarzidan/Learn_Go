package main
import "fmt"

const NMAX = 1000

type wisudawan struct {
    nama, nim        string
    eprt, semester, ipk float64
}

func main() {
    var wisudawans [NMAX]wisudawan
    var total int

    total = isiWisudawan(&wisudawans)

    fmt.Println(eprtTertinggi(wisudawans, total))
    fmt.Println(ipkTerendah(wisudawans, total))
    fmt.Println(rataSemester(wisudawans, total))
}

func isiWisudawan(wisudawans *[NMAX]wisudawan) int {
    var total int
    var nim string

    for i := 0; i < NMAX; i++ {
        fmt.Scan(&wisudawans[i].nama, &nim, &wisudawans[i].eprt, &wisudawans[i].semester, &wisudawans[i].ipk)
        if nim == "none" {
            total = i
            break
        }
    }

    return total
}

func eprtTertinggi(wisudawans [NMAX]wisudawan, total int) float64 {
    var max float64

    max = wisudawans[0].eprt
    for i := 1; i < total; i++ {
        if wisudawans[i].eprt > max {
            max = wisudawans[i].eprt
        }
    }

    return max
}

func ipkTerendah(wisudawans [NMAX]wisudawan, total int) float64 {
    var min float64

    min = wisudawans[0].ipk
    for i := 1; i < total; i++ {
        if wisudawans[i].ipk < min {
            min = wisudawans[i].ipk
        }
    }

    return min
}

func rataSemester(wisudawans [NMAX]wisudawan, total int) float64 {
    var sum float64

    for i := 0; i < total; i++ {
        sum += wisudawans[i].semester
    }

    return sum / float64(total)
}