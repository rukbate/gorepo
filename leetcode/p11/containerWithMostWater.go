package main

import "fmt"

func maxArea(height []int) int {
	maxA := 0
	i, j := 0, len(height) - 1

	for i < j {
		ca := 0
		if height[i] < height[j] {
			ca = height[i] * (j - i)
			i++
		} else {
			ca = height[j] * (j - i)
			j--
		}

		if ca > maxA {
			maxA = ca
		}
	}

	return maxA
}

func main() {
	arr := []int {1,8,6,2,5,4,8,3,7}

	fmt.Println(maxArea(arr))
}