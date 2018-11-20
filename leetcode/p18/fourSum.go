package main
import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	
	sort.Ints(nums)
	length := len(nums)

	if length < 4 {
		return res
	}

	for i := 0; i < length - 3; i++ {
		if target >= 0 && nums[i] + nums[i + 1] > target {
			break 
		}
		for j := i + 1; j < length - 2; j++ {
			m, n := j + 1, length - 1
			for m < n {
				if nums[i] + nums[j] + nums[m] + nums[n] < target {
					m++
				} else if nums[i] + nums[j] + nums[m] + nums[n] > target {
					n--
				} else {
					existing := false
					for x := 0; x < len(res); x++ {
						if nums[i] == res[x][0] && nums[j] == res[x][1] && nums[m] == res[x][2] {
							existing = true
						}
					}
					
					if !existing {
						res = append(res, []int {nums[i], nums[j], nums[m], nums[n]})
					}
					m++
					n--
				}
			}
		}
	}

	return res
}

func main() {
	nums := []int {-5,-4,-3, -2, 1, 3,3,5}
	target := -11

	fmt.Println(fourSum(nums, target))
}