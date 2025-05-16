package main

import "fmt"

func main() {

	//floor of an array number
	//find the Greatest number smaller than target

	nums := []int{1, 3, 6, 7, 12, 16, 23, 59}
	target := 15

	fmt.Println(FloorBinarySearch(nums, target))
	//fmt.Println(CeilBinarySearch(nums, target))
}

//return the index of the target of the largest of smaller than target
func FloorBinarySearch(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for end >= start {
		mid := (start + end) / 2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}

	}
	return end
}
