package main

import "fmt"

func main() {
	arr := []int{4, 2, 0, 2, 2, 1, 4, 4, 1, 4, 3, 2}
	removeElement(arr, 4)
	fmt.Println(arr)
	//arr2 := []int{3, 3}
	//removeElement(arr2, 3)
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}
