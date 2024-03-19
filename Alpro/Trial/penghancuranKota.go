package main

import "fmt"

func main() {
	var banyakMusuh, kekuatanKenshin, kecepatanKenshin, totalMusuhKalah int
	fmt.Scan(&banyakMusuh)
	fmt.Scan(&kekuatanKenshin, &kecepatanKenshin)

	totalMusuhKalah = 0

	for i := 0; i < banyakMusuh; i++ {
		var kekuatanShishio, kecepatanShishio int
		fmt.Scan(&kekuatanShishio, &kecepatanShishio)

		if kekuatanKenshin >= 3 && kecepatanKenshin >= 3 {
			if kekuatanShishio < kekuatanKenshin && kecepatanShishio < kecepatanKenshin {
				if kekuatanKenshin > kecepatanKenshin {
					kekuatanKenshin = kekuatanKenshin - 6
				} else {
					kecepatanKenshin = kecepatanKenshin - 6
				}
				totalMusuhKalah++
			} else if kekuatanShishio >= kekuatanKenshin && kecepatanShishio < kecepatanKenshin {
				kecepatanKenshin = kecepatanKenshin - 6
				totalMusuhKalah++
			} else if kekuatanShishio < kekuatanKenshin && kecepatanShishio >= kecepatanKenshin {
				kekuatanKenshin = kekuatanKenshin - 6
				totalMusuhKalah++
			}
		}
	}
	fmt.Println(totalMusuhKalah, kekuatanKenshin, kecepatanKenshin)
}
