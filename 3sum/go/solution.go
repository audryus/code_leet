package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			ln := nums[left]
			rn := nums[right]

			sum := nums[i] + ln + rn

			if sum < 0 {
				left = left + 1
			} else if sum > 0 {
				right = right - 1
			} else {
				result = append(result, []int{nums[i], ln, rn})

				left = left + 1

				for nums[left] == nums[left-1] && left < right {
					left = left + 1
				}
			}
		}
	}

	return result
}
