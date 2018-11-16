package main

import "fmt"

func maxArea(height []int) int {
	maxH := 0
	maxA := 0
	for i := 0; i < len(height); i++ {
		if height[i] > maxH {
			maxH = height[i]
		}
	}
	
	for step := len(height) - 1; step > 0; step-- {
		if step >= maxA / maxH {
			for i := 0; i + step < len(height); i++ {
				var ca int
				if height[i] > height[i + step] {
					ca = step * height[i + step]
				} else {
					ca = step * height[i]
				}

				if ca > maxA {
					maxA = ca
				}
			}
		}
	}

	return maxA
}

func main() {
	arr := []int {1,8, 9}

	fmt.Println(maxArea(arr))
}