package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))

}
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sLen := len(s)
	tLen := len(t)

	if sLen > tLen {
		return false
	}

	j := 0
	for i := range tLen {
		if t[i] == s[j] {
			if j == sLen-1 {
				return true
			}
			j += 1
		}
	}

	return false
}
