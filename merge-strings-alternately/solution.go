package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
func mergeAlternately(word1 string, word2 string) string {
	A, B := len(word1), len(word2)
	a, b := 0, 0
	s := make([]string, (A + B))

	word := 1

	for a < A && b < B {
		if word == 1 {
			s = append(s, string(word1[a]))
			a++
			word = 2
		} else {
			s = append(s, string(word2[b]))
			b++
			word = 1
		}
	}

	for a < A {
		s = append(s, string(word1[a]))
		a++
	}
	for b < B {
		s = append(s, string(word2[b]))
		b++
	}

	return strings.Join(s, "")
}
