package main

import "fmt"

func main() {
	fmt.Printf("%v\n", summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Printf("%v\n", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
func summaryRanges(nums []int) []string {
	var ans []string
	i := 0
	size := len(nums)
	for i < size {
		start := nums[i]
		for i < size-1 && nums[i]+1 == nums[i+1] {
			i += 1
		}

		if start != nums[i] {
			ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			ans = append(ans, fmt.Sprintf("%d", nums[i]))
		}
		i += 1
	}

	return ans
}
