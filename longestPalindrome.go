package main;

import "fmt"

func longestPalindrome(s string) string {
	if(len(s) == 0) {
		return "";
	}

	m := 0; n := 0; max := 0;
	for i := 0; i < len(s); i++ {
		j := 1;
		for {
			if(i + j >= len(s) || i - j < 0) {
				break;
			}

			if(s[i - j] == s[i + j]) {
				if(j * 2 + 1 >= max) {
					max = j * 2 + 1;
					m = i - j;
					n = i + j;
				}

				j++;
			} else {
				break;
			}
		}

		j = 1;
		for {
			if(i + j >= len(s) || i - j + 1 < 0) {
				break;
			}

			if(s[i - j + 1] == s[i + j]) {
				if(j * 2 >= max) {
					max = j * 2;
					m = i - j + 1;
					n = i + j;
				}

				j++;
			} else {
				break;
			}
		}
	}

    return s[m:n+1];
}

func main() {
	fmt.Printf("%s\n", longestPalindrome("ababc"));
}