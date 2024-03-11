package main
import "fmt"

func main() {
	var koinLumin, koinQuantumShard, koinGalactum, koinNebulaCrown, koinStellaris int
    fmt.Scan(&koinLumin)

    // 1 QuantumShard = 10 Galactum
    // 1 Galactum = 3 NebulaCrown
    // 1 NebulaCrown = 2 Stellaris
    // 1 Stellaris = 20 Lumin

    koinQuantumShard = koinLumin / (20 * 2 * 3 * 10) // 1381 / 1200 = 1
    koinLumin = koinLumin % (20 * 2 * 3 * 10) // 1381 % 1200 = 181

    koinGalactum = koinLumin / (20 * 2 * 3) // 181 / 120 = 1
    koinLumin = koinLumin %  (20 * 2 * 3) // 181 % 120 = 61

    koinNebulaCrown = koinLumin / (20 * 2) // 61 / 40 = 1
    koinLumin = koinLumin %  (20 * 2) // 61 % 40 = 21

    koinStellaris = koinLumin / 20 // 21 / 20 = 1
    koinLumin = koinLumin % 20 // 21 % 20 = 1

    fmt.Println(koinQuantumShard, koinGalactum, koinNebulaCrown, koinStellaris, koinLumin)
}



// func main() {
//     var kLumin, kQuantumShard, kGalactum, kNebulaCrown, kStellaris int

//     fmt.Scan(&kLumin)

//     kQuantumShard = kLumin / 1200 
//     kLumin = kLumin - (kQuantumShard * 1200) 

//     kGalactum = kLumin / 120 
//     kLumin = kLumin - (kGalactum * 120) 

//     kNebulaCrown = kLumin / 40 
//     kLumin = kLumin - (kNebulaCrown * 40) 

//     kStellaris = kLumin / 20 
//     kLumin = kLumin - (kStellaris * 20)

//     fmt.Println(kQuantumShard, kGalactum, kNebulaCrown, kStellaris, kLumin)
// }



// var koinLumin, koinQuantumShard, koinGalactum, koinNebulaCrown, koinStellaris int
// fmt.Scan(&koinLumin)

// koinQuantumShard = koinLumin / (20 * 2 * 3 * 10)
// koinLumin = koinLumin % (20 * 2 * 3 * 10)

// koinGalactum = koinLumin / (20 * 2 * 3)
// koinLumin = koinLumin % (20 * 2 * 3)

// koinNebulaCrown = koinLumin / (20 * 2)
// koinLumin = koinLumin % (20 * 2)

// koinStellaris = koinLumin / 20
// koinLumin = koinLumin % 20

// fmt.Println(koinQuantumShard, koinGalactum, koinNebulaCrown, koinStellaris, koinLumin)

// var koinLumin, koinQuantumShard, koinGalactum, koinNebulaCrown, koinStellaris int
// fmt.Scan(&koinLumin)

// koinQuantumShard = koinLumin / (20 * 2 * 3 * 10)
// koinLumin = koinLumin % (20 * 2 * 3 * 10)

// koinGalactum = koinLumin / (20 * 2 * 3)
// koinLumin = koinLumin % (20 * 2 * 3)

// koinNebulaCrown = koinLumin / (20 * 2)
// koinLumin = koinLumin % (20 * 2)

// koinStellaris = koinLumin / 20
// koinLumin = koinLumin % 20

// fmt.Println(koinQuantumShard, koinGalactum, koinNebulaCrown, koinStellaris, koinLumin)
