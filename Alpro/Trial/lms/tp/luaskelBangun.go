package main

import "fmt"

func main() {
	var r, s, ll, lp, kl, kp, tl, tp float64

	fmt.Scan(&r, &s)
	if r != 0 && s != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	}
	for r != 0 && s != 0 {
		hitungLuasKelilingLingkaran(r, &ll, &kl)
		hitungLuasKelilingPersegi(s, &lp, &kp)
		hitungTotal(ll, lp, kl, kp, &tl, &tp)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, ll, lp, kl, kp, tl, tp)
		fmt.Scan(&r, &s)
	}
}

func hitungLuasKelilingLingkaran(r float64, ll, kl *float64) {
	*ll = 3.14 * r * r
	*kl = 2 * 3.14 * r
}

func hitungLuasKelilingPersegi(s float64, lp, kp *float64) {
	*lp = s * s
	*kp = 4 * s
}

func hitungTotal(ll, lp, kl, kp float64, tl, tp *float64) {
	*tl = ll + lp
	*tp = kl + kp
}
