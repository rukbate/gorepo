package main
import (
	"fmt"
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	sum := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {		
		j, k := i + 1, len(nums) - 1
		for j < k {
			if nums[i] + nums[j] + nums[k] > target {
				if math.Abs(float64(nums[i] + nums[j] + nums[k] - target)) < math.Abs(float64(sum - target)) {
					sum = nums[i] + nums[j] + nums[k]
				}
				k--
			} else if nums[i] + nums[j] + nums[k] < target {
				if math.Abs(float64(nums[i] + nums[j] + nums[k] - target)) < math.Abs(float64(sum - target)) {
					sum = nums[i] + nums[j] + nums[k]
				}
				j++
			} else {
				sum = target
				break
			}
		}
	}

	return sum	
}

func main() {
	nums := []int {-5, -1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}