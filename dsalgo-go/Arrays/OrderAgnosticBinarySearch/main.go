package main

import "fmt"

func main() {
	elements := []int{12, 31, 45, 67, 78, 89, 90, 112, 145}
	reversElements := []int{150, 130, 90, 89, 76, 46, 32, 11, 1}
	target := 89

	fmt.Println(OrderAgnosticBinarySearch(elements, target))
	fmt.Println(OrderAgnosticBinarySearch(reversElements, target))

}

func OrderAgnosticBinarySearch(nums []int, target int) int {

	//note :compare the first and last number not first 2 it can be {3, 3, 5, 6,7, 89, 110}
	AscendingTrue := nums[0] < nums[len(nums)-1]
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if AscendingTrue {
			if nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] > target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

	}
	return -1
}
