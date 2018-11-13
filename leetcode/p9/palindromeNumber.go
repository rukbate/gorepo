package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x < 10 {
        return true
    } else if x < 100 {
        return x % 11 == 0
    } else {
		res := 0
		n := x
		for n > 0 {		
			r := n % 10;
			res = 10 * res + r;
			n = n / 10;
		}

		return res == x
    }    
}

func main() {
	fmt.Println(isPalindrome(10));
}