package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(arr)
	fmt.Println(arr)
	//arr2 := []int{3, 3}
	//removeElement(arr2, 3)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1

}
