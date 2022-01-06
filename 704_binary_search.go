package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}

	}
	return -1

}

func main() {
	var arr = [6]int{-1, 0, 3, 5, 9, 12}
	fmt.Print(search(arr[:], 9))
}
