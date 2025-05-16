package main

import "fmt"

func main() {

	//ceiling of an array
	// find the smallest of greater than target

	nums := []int{1, 3, 6, 7, 12, 16, 23, 59}
	target := 15

	fmt.Println(CeilingBinarySearch(nums, target))
}

//return index of the target or smallest of the greater than target
func CeilingBinarySearch(nums []int, target int) int {

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}
