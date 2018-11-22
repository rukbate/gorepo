package main
import "fmt"

func generateParenthesis(n int) []string {
	res := make(map[string]bool)
	cache := make(map[int][]string)
	
	return gp(n, res, cache)
}

func gp(n int, res map[string]bool, cache map[int][]string) []string {
	if r, found := cache[n]; found {
		return r
	}
	
	r := make([]string, 0)
	if n == 0 {
		
	} else if n == 1 {
		r = append(r, "()")		
	} else {
		for i := 1; i < n; i++ {
			p1 := gp(i, res, cache)
			p2 := gp(n - i, res, cache)

			for m := 0; m < len(p1); m++ {
				for n := 0; n < len(p2); n++ {
					s := []string {p1[m] + p2[n], p2[n] + p1[m]}
					for j := 0; j < 2; j++ {
						if !res[s[j]] {
							res[s[j]] = true
							r = append(r, s[j])
						}
					}
				}
			}

			if i == 1 {
				for n := 0; n < len(p2); n++ {
					s := "(" + p2[n] + ")"
					if !res[s] {
						res[s] = true
						r = append(r, s)
					}
				}
			}
		}
	}

	cache[n] = r
	return r
}

func main() {
	res := generateParenthesis(5)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Println(len(res))
}