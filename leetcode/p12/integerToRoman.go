package main

import "fmt"

func intToRoman(num int) string {
	integer := [13]int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := [13]string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	ct := 0
	buffer := make([]byte, 0, 50)

	for p := 0; p < 13; p++ {
		for num - ct >= integer[p] {
			ct += integer[p]
			buffer = append(buffer, roman[p][0])
			if len(roman[p]) == 2 {
				buffer = append(buffer, roman[p][1])
			}
		}
	}
	
	return string(buffer)
}

func main() {
	fmt.Println(intToRoman(3))
}