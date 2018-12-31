package main
import "fmt"

func nextPermutation(nums []int)  {
	i := len(nums) - 1
	for i > 0 {
		if nums[i] > nums[i - 1] {
			break
		}

		i--
	}

	if i > 0 {		
		k := i
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i - 1] && nums[j] <= nums[k] {
				k = j
			}			
		}

		swap(nums, i - 1, k)
	}

	reverse(nums, i)	
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start int) {
	i, j := start, len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func main() {
	nums := []int {2, 3, 1, 3, 3}
	nextPermutation(nums)

	fmt.Println(nums)
}