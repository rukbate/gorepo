package main

import "fmt"

func convert(s string, numRows int) string {
	if(len(s) == 0 || numRows <= 1) {
		return s;
	}
	queue := make([]int, 0);
	p := 0;
	step := 2 * (numRows - 1);
	set := make([]bool, len(s) + step);

	i := 0;
	for ; i < len(s); i += step {
		queue = append(queue, i);
		set[i] = true;
	}

	queue = append(queue, i);
	set[i] = true;

	var x int;
	for {
		if(len(queue) <= p) {
			break;
		}

		x = queue[p];
		p++;

		if(x - 1 >= 0) {
			if(set[x - 1] != true) {
				queue = append(queue, x - 1);
				set[x - 1] = true;
			}			
		}
		if(x + 1 < len(s)) {
			if(set[x + 1] != true) {
				queue = append(queue, x + 1);
				set[x + 1] = true;
			}
		}
	}

	j := 0;
	res := make([]byte, len(s));
	for i := 0; i < len(queue); i++ {
		if(queue[i] < len(s)) {
			res[j] = s[queue[i]];
			j++;
		}
	}
	return string(res);
}

func main() {
	fmt.Printf("%s\n", convert("PAYPALISHIRING", 3));
}