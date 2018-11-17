package main

import "fmt"

func romanToInt(s string) int {
    integer := [13]int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := [13]string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	
	i := 0
	res := 0

	for p := 0; p < 13; p++ {
		n := len(roman[p])

		for i + n <= len(s) && s[i: i + n] == roman[p] {
			res += integer[p]
			i = i + n
		}
	}

	return res
}

func main() {
	roman := "III"

	fmt.Println(romanToInt(roman))
}