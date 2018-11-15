package main

import "fmt"

func isMatch(s string, p string) bool {
	i := 0
	j := 0
	
	return match(s, p, i, j)
}

func match(s string, p string, i int, j int) bool {
	if i == len(s) && j == len(p) {
		return true
	} else if i < len(s) && j < len(p) {
		if j + 1 < len(p) && p[j + 1] != '*' {
			return matchByte(s[i], p[j]) && match(s, p, i + 1, j + 1)
		} else if j + 1 < len(p) && p[j + 1] == '*' {
			return (matchByte(s[i], p[j]) && match(s, p, i + 1, j + 2)) || (matchByte(s[i], p[j]) && match(s, p, i + 1, j)) || match(s, p, i, j + 2)
		} else {
			return i + 1 == len(s) && matchByte(s[i], p[j])
		}
	} else if j < len(p) {
		a := 0
		b := 0
		for j < len(p) {
			if p[j] != '*' {
				a++
			} else {
				b++
			}

			j++
		}
		return a == b
	} else {
		return false
	}
}

func matchByte(a byte, b byte) bool {
	return a == b || b == '.'
}

func main() {
	s := "a"
	p := "ab*"

	fmt.Println(isMatch(s, p))
}