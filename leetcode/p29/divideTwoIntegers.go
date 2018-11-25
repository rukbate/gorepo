package main
import "fmt"

func divide(dividend int, divisor int) int {
	sdd, dd := signAndAbs(dividend)
	sdr, dr := signAndAbs(divisor)

	sr := sdd == sdr

	if dividend >= 2147483647 {
		if divisor == 1 {
			return 2147483647
		} else if divisor == -1 {
			return -2147483647
		} 
	} else if dividend <= -2147483648 {
		if divisor == 1 {
			return -2147483648
		} else if divisor == -1 {
			return 2147483647
		}
	}
	
	n := 0
	add := dr
	for add <= dd {
		add += dr
		n++
	}

	if !sr {
		n = 0 - n
	}

	return n
}

func signAndAbs(num int) (bool, int) {
	if num >= 0 {
		return true, num
	} 
	
	return false, 0 - num	
}

func main() {
	fmt.Println(divide(-2147483648, 2))
}