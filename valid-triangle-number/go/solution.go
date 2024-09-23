package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4})) // 3
	fmt.Println(triangleNumber([]int{4, 2, 3, 4})) // 4
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := len(nums) - 1; i > 1; i-- {
		left := 0
		right := i - 1

		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count = count + (right - left)
				right = right - 1
			} else {
				left = left + 1
			}
		}
	}

	return count
}
