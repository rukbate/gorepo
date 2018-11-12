package main

import "fmt"

func myAtoi(str string) int {
	n := 0
	i := 0
	sign := 1

	for i < len(str) && str[i] == ' ' {		
		break		
	}

	if i < len(str) && str[i] == '-' {
		i++
		sign = -1
	} else if i < len(str) && str[i] == '+' {
		i++
	} 
	
	for i < len(str) && int(str[i]) >= 48 && int(str[i]) <= 57 {
		n = n * 10 + int(str[i] - 48)
		i++

		if n < -2147483648 || n > 2147483647 {
			break
		}	
	}
	
	n = n * sign

	if n < -2147483648 || n > 2147483647 {
		if sign == -1 {
			return -2147483648
		} 
			
		return 2147483647
	} 
	
	return n
}

func main() {
	fmt.Println(myAtoi("3459"))
}