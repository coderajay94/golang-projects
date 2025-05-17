package main

import "fmt"

//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
func main() {

	nums := []int{1, 2, 4, 4, 4, 4, 7, 8, 12, 14, 34, 45}
	target := 4

	position := SearchRange(nums, target)

	fmt.Println("starting from ", position[0], " and ending at", position[1])
}

//returns starting and end index of target
func SearchRange(nums []int, target int) []int {

	resp := []int{-1, -1}
	resp[0] = SearchLeft(nums, target, true)
	resp[1] = SearchLeft(nums, target, false)

	return resp
}

//pass extra bool in binary search to check if wanna check the starting point or ending point
//need to call this function twice, as 2(log n) = log n complexity
func SearchLeft(nums []int, target int, isLeft bool) int {

	start, end := 0, len(nums)-1
	ans := -1
	for start <= end {
		mid := (start + end) / 2

		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			//potential answer but keep moving to left
			ans = mid

			//keep shortening the path on the left and right side till loop breaks
			if isLeft {
				end = mid - 1
			} else {
				start = mid + 1
			}

		}

	}
	return ans
}
