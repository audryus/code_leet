package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int) {
	nextNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nextNonZero], nums[i] = nums[i], nums[nextNonZero]
			nextNonZero++
		}
	}
}
