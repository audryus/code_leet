package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxArea := 0

	for left < right {
		w := right - left
		h := min(height[left], height[right])
		area := w * h

		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left = left + 1
		} else {
			right = right - 1
		}

	}

	return maxArea
}
