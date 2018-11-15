package main

import "fmt"

func isMatch(s string, p string) bool {
	i := 0
	j := 0

	arr := make([][]int, len(s), len(s))
	for x := 0; x < len(s); x++ {
		brr := make([]int, len(p), len(p))
		arr[x] = brr
	}
	
	return match(s, p, i, j, arr)
}

func match(s string, p string, i int, j int, arr [][]int) bool {
	if i < len(s) && j < len(p) && arr[i][j] != 0 {
		return arr[i][j] == 1
	} 

	n, k := nextUnit(p, j)
	y := false
	if i >= len(s) {
		y = (k == -1 || (len(n) == 2 && match(s, p, i, k, arr)))
	} else if len(n) == 1 {
		y = matchUnit(s[i:i+1], n) && match(s, p, i + 1, k, arr)
	} else if len(n) == 2 {
		y = match(s, p, i, k, arr) || (matchUnit(s[i: i + 1], n) && (match(s, p, i + 1, k, arr) || match(s, p, i + 1, k - 2, arr)))
	} else {
		y = i >= len(s)
	}

	if( i < len(s) && j < len(p)) {
		if y {
			arr[i][j] = 1
		} else {
			arr[i][j] = -1
		}
	}

	return y
}

func nextUnit(s string, i int) (string, int) {
	if i < len(s) {
		if i + 1 < len(s) && s[i + 1] == '*' {
			return s[i: i + 2], i + 2
		} 
	
		return (string)(s[i]), i + 1		
	}
	
	return "", -1	
}

func matchUnit(s string, p string) bool {
	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	} else if len(p) == 2 {
		return len(s) == 0 || (s[0] == p[0] || p[0] == '.')
	} else {
		return false
	}
}

func main() {
	s := "ab"
	p := "d*c*a*b*"

	fmt.Println(isMatch(s, p))
}