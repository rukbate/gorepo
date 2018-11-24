package main
import "fmt"

func removeDuplicates(nums []int) int {
	if(len(nums) < 2) {
		return len(nums)
	}
	
	p := nums[0]
	i := 1

	for i < len(nums) {
		if nums[i] == p {
			nums = append(nums[:i], nums[i + 1:]...)
		} else {
			p = nums[i]
			i++
		}
	}
	
	return len(nums)
}

func main() {
	nums := []int {0,0,1,1,1,2,2,3,3,4}

	rem := removeDuplicates(nums)

	fmt.Printf("%d:\n", rem)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}