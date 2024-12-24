package main

import "fmt"

func main() {
	fmt.Printf("%v\n", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
func productExceptSelf(nums []int) []int {
	leftMult := 1
	rightMult := 1
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)

	for i := range n {
		j := n - i - 1
		left[i] = leftMult
		right[j] = rightMult

		leftMult *= nums[i]
		rightMult *= nums[j]
	}

	var ans []int
	for i := 0; i < n; i++ {
		ans = append(ans, left[i]*right[i])
	}

	return ans
}
