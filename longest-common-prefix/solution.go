package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	l := len(strs)

	for i := range strs[0] {
		if strs[0][i] != strs[l-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]

}
