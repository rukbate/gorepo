package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Printf("%s\n", "FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Printf("%s\n", "Fizz")
		} else if i % 5 == 0 {
			fmt.Printf("%s\n", "Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
