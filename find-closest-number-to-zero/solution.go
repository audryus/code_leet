package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8}))
	fmt.Println(findClosestNumber([]int{2, -1, 1}))
}

func findClosestNumber(nums []int) int {
	closest := nums[0]

	for i := 1; i < len(nums); i++ {
		if Abs(nums[i]) < Abs(closest) {
			closest = nums[i]
		}
	}

	if closest < 0 && slices.Contains(nums, Abs(closest)) {
		return Abs(closest)
	}

	return closest
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
