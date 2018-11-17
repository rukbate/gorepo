package main
import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	base := strs[0]
	pos := 0
	end := false
	for i := 0; i <= len(base) && !end; i++ {
		prefix := base[:i]

		for j := 1; j < len(strs); j++ {
			str := strs[j]
			if i > len(str) || str[:i] != prefix {
				end = true
				break
			} 
		}

		if !end {
			pos = i
		}
	}	

	return base[:pos]
}

func main() {
	strs := []string {"c", "c"}
	fmt.Println(longestCommonPrefix(strs))
}