package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(arr)
	fmt.Println(arr)

	arr = []int{4, 2, 0, 3, 2, 5}
	trap(arr)
	fmt.Println(arr)
}

func trap(height []int) int {

}
