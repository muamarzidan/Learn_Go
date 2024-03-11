package main

import "fmt"

// func main() {
//     var celciusAwal, celciusAkhir, step float64
//     fmt.Scanf("%f %f %f", &celciusAwal, &celciusAkhir, &step)

//     fmt.Printf("%10s %10s %10s %10s\n", "Celsius", "Reamur", "Fahrenheit", "Kelvin")

//     for i := celciusAwal; i <= celciusAkhir; i += step {
//         reamur := convertToReamur(i)
//         fahrenheit := convertToFahrenheit(i)
//         kelvin := convertToKelvin(i)

//         fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", i, reamur, fahrenheit, kelvin)
//     }
// }


func main() {
    var celciusAwal, celciusAkhir, step float64
    fmt.Scanf("%f %f %f", &celciusAwal, &celciusAkhir, &step)

    fmt.Printf("%10s %10s %10s %10s\n", "Celsius", "Reamur", "Fahrenheit", "Kelvin")

    for i := celciusAwal; i <= celciusAkhir; i += step {
        reamur := convertToReamur(i)
        fahrenheit := convertToFahrenheit(i)
        kelvin := convertToKelvin(i)

        fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", i, reamur, fahrenheit, kelvin)
    }
}

func convertToReamur(celsius float64) float64 {
    return 4 * celsius / 5
}

func convertToFahrenheit(celsius float64) float64 {
    return (9 * celsius / 5) + 32
}

func convertToKelvin(celsius float64) float64 {
    return celsius + 273.00
}

