package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
func romanToInt(s string) int {
	d := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	summ := 0
	n := len(s)

	i := 0

	for i < n {
		if i < n-1 && d[s[i]] < d[s[i+1]] {
			summ += d[s[i+1]] - d[s[i]]
			i += 2
		} else {
			summ += d[s[i]]
			i += 1
		}
	}
	return summ
}
