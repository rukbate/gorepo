package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("A")
		}
		fmt.Print("\n")
	}
}
