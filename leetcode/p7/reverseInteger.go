package main

import "fmt"

func reverse(x int) int {
	n := x;

	res := 0;
	minus := 1;
	if(x < 0) {
		minus = -1;
		n = 0 - x;
	}

	for {
		r := n % 10;
		res = 10 * res + r;
		n = n / 10;

		if(n <= 0) {
			break;
		}
	}

	if(res > 2147483648 || res <= -2147483648) {
		return 0;
	}
	
	return res * minus;
}

func main() {
	fmt.Printf("%d\n", reverse(43567));
}