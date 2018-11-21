package main
import "fmt"

func isValid(s string) bool {
	paren := map[byte]byte {
		')': '(',
		']': '[',
		'}': '{',
	}

	length := len(s)
	stack := make([]byte, length)

	p := -1
	err := false
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(', '[', '{':
			p++
			stack[p] = s[i]			
		case ')', ']', '}':
			if p >= 0 && stack[p] == paren[s[i]] {
				p--
			} else {
				err = true
				break
			}
		}
	}

	return !err && p == -1
}

func main() {
	fmt.Println(isValid("({)[}]"))
}