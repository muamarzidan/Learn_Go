package main

import "fmt"

func main() {
    var r, t float64
    fmt.Scan(&r, &t)
    fmt.Printf("%f", volumeKerucut(r, t))
}

func volumeKerucut(r, t float64) float64 {
    const phi float64 = 3.14
    var volume float64
    volume = phi * r * r * t / 3
    return volume
}

