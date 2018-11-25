package main
import "fmt"

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	
	if hl < nl {
		return -1
	}
    if nl == 0 {
		return 0
	}

	for i := 0; i <= hl - nl; i++ {
		match := true
		j, k := i, 0
		for j < hl && k < nl {
			if haystack[j] != needle[k] {
				match = false
				break
			} else {
				j++
				k++
			}
		}

		if match {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("", ""))
}