package main
import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	cache := make(map[int]int)

	mid := (n1 + n2) / 2
	if (n1 + n2) % 2 == 1 {
		cache[mid] = 0
	} else {
		mid := (n1 + n2) / 2
		cache[mid - 1] = 0
		cache[mid] = 0
	}

	i, j := 0, 0
	c := -1
	for i < n1 && j < n2 {
		c++
		_, ok := cache[c]
		if nums1[i] < nums2[j] {
			if ok {
				cache[c] = nums1[i]
			}
			i++
		} else {
			if ok {
				cache[c] = nums2[j]
			}
			j++
		}
	}

	for i < n1 {
		c++
		if _, ok := cache[c]; ok {
			cache[c] = nums1[i]			
		}
		i++
	}

	for j < n2 {
		c++
		if _, ok := cache[c]; ok {
			cache[c] = nums2[j]
		}
		j++
	}

	sum := 0
	for _, val := range cache {
		sum += val
	}

	return float64(sum) / float64(len(cache))
}

func main() {
	nums1 := []int {1, 3, 4, 7, 9}
	nums2 := []int {2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}