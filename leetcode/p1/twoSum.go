package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	
	for i := 0; i < len(nums); i++ {
		if val, ok := cache[target - nums[i]]; ok {		
			return []int {val, i};
		}

		cache[nums[i]] = i
	}

	return nil
}

func main() {
	nums := []int {2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}