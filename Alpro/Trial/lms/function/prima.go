func main() {
	var input int

	fmt.Scan(&input)

	fmt.Print(prima(input))
}

func prima(input int) string {
	if input < 2 {
		return "Bukan Prima"
	} else if input == 2 {
		return "Prima"
	} else {
		if input / input == 1 && input / 1 == input  && input % 2 != 0 {
			return "Prima"
		} else {
			return "Bukan Prima"
		}
	}
}