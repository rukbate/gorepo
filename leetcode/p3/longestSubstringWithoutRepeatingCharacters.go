package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	res := 0

	cache := make(map[byte]int)
	for i, j := 0, 0; j < n; j++ {
		if c, ok := cache[s[j]]; ok {
			if c > i {
				i = c
			}
		}

		if j - i + 1 > res {
			res = j - i + 1
		}
		cache[s[j]] = j + 1
	}

	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}