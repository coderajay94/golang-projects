package main

import "fmt"

//https://leetcode.com/problems/find-peak-element
func main() {

	//nums2 := []int{1, 2, 4, 5, 6, 8, 7, 4, 3, 1}
	nums := []int{1, 2, 3, 1}

	fmt.Println("peak element  index is :", findPeakElement(nums))
}

//using bineary search logic
//returns index
func findPeakElement(nums []int) int {

	start, end := 0, len(nums)-1

	for start < end {
		mid := (start + end) / 2

		//logic is when mid+1 element is greater than mid
		// 5, 6
		// it can't be the peak so move mid+1
		//
		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
		//as in 5, 6, 4
		//when the order is decreasing chances are this is the potential answer
		//still narrow it down by making it end = mid
	}

	return start
	//return nums[start]
}
