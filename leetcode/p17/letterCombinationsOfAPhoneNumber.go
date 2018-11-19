package main
import "fmt"

func letterCombinations(digits string) []string {
	letters := []string {"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	
	return lc(digits, &letters)
}

func lc(digits string, letters *[]string) []string {
	res := make([]string, 0)

	if len(digits) == 0 {
		return res
	} else if len(digits) == 1 {
		str := (*letters)[int(digits[0] - 48) - 1]
		for i := 0; i < len(str); i++ {
			res = append(res, string(str[i]))
		}

		return res
	} else {
		str := (*letters)[int(digits[0] - 48) - 1]
		subres := lc(digits[1:], letters)
		
		for i := 0; i < len(str); i++ {
			for j := 0; j < len(subres); j++ {
				res = append(res, string(str[i]) + subres[j])
			}			
		}

		return res
	}
}

func main() {
	fmt.Println(letterCombinations("239"))
}